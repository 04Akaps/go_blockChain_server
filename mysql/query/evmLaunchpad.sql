-- name: CreateEvmLaunchpad :execresult
INSERT INTO evmLaunchpad (
    eoa_address,
    ca_address,
    chain_id,
    created_at
) VALUES (
   ?, ?, ?, ?
);

-- name: GetMyLaunchpad :one
SELECT * FROM evmLaunchpad
WHERE eoa_address = ? LIMIT 1;

-- name: GetMyAllLaunchpad :many
SELECT * FROM evmLaunchpad;