CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL, -- keep track of the client type
  "client_ip" varchar NOT NULL, -- where the client is connecting to the server, client ip address
  "is_blocked" boolean NOT NULL DEFAULT false, -- block the session if the refresh token is compromised
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
