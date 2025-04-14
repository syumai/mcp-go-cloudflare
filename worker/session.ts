import { DurableObject } from "cloudflare:workers";

type SSESession = {
  sessionId: string;
};

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

  async store(sessionId: string, session: SSESession): Promise<void> {
    this.sql.exec("INSERT INTO sessions (session_id) VALUES ($1)", [sessionId]);
  }

  async range(
    f: (sessionId: string, session: SSESession) => boolean
  ): Promise<void> {
    const rows = await this.sql.exec("SELECT * FROM sessions");
    for (const row of rows) {
      const session: SSESession = { sessionId: row.session_id };
      if (!f(row.session_id, session)) {
        break;
      }
    }
  }

  async load(sessionId: string): Promise<SSESession | null> {
    const row = await this.sql.exec(
      "SELECT * FROM sessions WHERE session_id = $1",
      [sessionId]
    );
    if (row.length === 0) {
      return null;
    }
    return { sessionId: row[0].session_id };
  }

  async delete(sessionId: string): Promise<void> {
    this.sql.exec("DELETE FROM sessions WHERE session_id = $1", [sessionId]);
  }
}
