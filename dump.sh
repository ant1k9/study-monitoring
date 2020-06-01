#!/bin/bash
set -euo pipefail

cd /tmp

DB_BACKUP_PATH="study.dump"
DB_BACKUP_ARCHIVE="$DB_BACKUP_PATH.zip"
DROPBOX_PATH="/study/$(date +'%Y%m%d_%H%M%S').dump.zip"

# Make backup
pg_dump --data-only "$PG_CONNECT" > "$DB_BACKUP_PATH"
zip "$DB_BACKUP_ARCHIVE" "$DB_BACKUP_PATH"
rm "$DB_BACKUP_PATH"

curl -X POST https://content.dropboxapi.com/2/files/upload \
    --header "Authorization: Bearer $DB_BACKUP_TOKEN" \
    --header "Content-Type: application/octet-stream" \
    --header "Dropbox-API-Arg: {\"path\": \"$DROPBOX_PATH\",\"mode\": \"add\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false}" \
    --data-binary @"$DB_BACKUP_ARCHIVE"

rm "$DB_BACKUP_ARCHIVE"
