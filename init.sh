#!/bin/bash

# Check if project name is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <project-name>"
  exit 1
fi

PROJECT_NAME=$1
OLD_PROJECT_NAME="go-template"
OLD_IMPORT_PATH="github.com/thoughtgears/$OLD_PROJECT_NAME"
NEW_IMPORT_PATH="github.com/thoughtgears/$PROJECT_NAME"

# Update go.mod files
find . -name "go.mod" -exec sed -i '' "s|$OLD_IMPORT_PATH|$NEW_IMPORT_PATH|g" {} +

# Update import paths in .go files
find . -name "*.go" -exec sed -i '' "s|$OLD_IMPORT_PATH|$NEW_IMPORT_PATH|g" {} +

echo "Project name updated to $PROJECT_NAME"
