import { DurableObject } from "cloudflare:workers";

type Env = unknown;

export class SessionStore extends DurableObject {
  sql: SqlStorage;
  constructor(ctx: DurableObjectState, env: Env) {
    super(ctx, env);
    this.sql = ctx.storage.sql;

    this.sql.exec(`CREATE TABLE IF NOT EXISTS sessions(
      session_id    TEXT PRIMARY KEY
    );`);
  }

  async store(sessionId: string): Promise<void> {
    this.sql.exec("INSERT INTO sessions (session_id) VALUES ($1)", [sessionId]);
  }

  async list(): Promise<string[]> {
    const cursor = this.sql.exec("SELECT * FROM sessions");
    const rows = cursor.toArray();
    return rows.map((row) => row.session_id as string);
  }

  async load(sessionId: string): Promise<string | null> {
    const cursor = this.sql.exec(
      "SELECT * FROM sessions WHERE session_id = $1",
      [sessionId]
    );
    const rows = cursor.toArray();
    if (rows[0] === undefined) {
      return null;
    }
    return rows[0].session_id as string;
  }

  async delete(sessionId: string): Promise<void> {
    this.sql.exec("DELETE FROM sessions WHERE session_id = $1", [sessionId]);
  }
}
