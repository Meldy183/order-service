CREATE TABLE IF NOT EXISTS "order" (
    id VARCHAR(36) PRIMARY KEY,
    item TEXT NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_order_item ON "order" (item);
