#!/bin/bash

# API Endpoints Testing Script
# This script contains curl instructions to test all endpoints of the TO-DO API application
# 
# Prerequisites:
# - The application should be running (default: http://localhost:8080)
# - curl should be installed
# - Modify the BASE_URL if the application is running on a different host/port

# Configuration
BASE_URL="http://localhost:8080"
API_VERSION="/api/v1"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}    TO-DO API - CURL Requests Script   ${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "Base URL: ${YELLOW}${BASE_URL}${NC}"
echo ""

# ============================================================================
# HEALTH CHECK ENDPOINT
# ============================================================================
echo -e "${GREEN}1. HEALTH CHECK ENDPOINT${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /health"
echo -e "${YELLOW}Description:${NC} Check if the application is running"
echo ""
echo "curl -X GET \"${BASE_URL}/health\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""
echo -e "${YELLOW}Alternative (with verbose output):${NC}"
echo "curl -v -X GET \"${BASE_URL}/health\""
echo ""
echo "---"
echo ""

# ============================================================================
# TASK ENDPOINTS
# ============================================================================
echo -e "${GREEN}2. TASK ENDPOINTS${NC}"
echo ""

echo -e "${YELLOW}2.1 Create a Task${NC}"
echo -e "${YELLOW}Endpoint:${NC} POST /api/v1/todo"
echo "curl -X POST \"${BASE_URL}${API_VERSION}/todo\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"title\": \"Sample Task\","
echo "    \"description\": \"This is a sample task\","
echo "    \"status_id\": 1,"
echo "    \"type_id\": 1,"
echo "    \"responsible_user_id\": 1,"
echo "    \"author_user_id\": 1,"
echo "    \"workflow_id\": 1,"
echo "    \"due_date\": \"2025-03-15T10:00:00Z\""
echo "  }'"
echo ""

echo -e "${YELLOW}2.2 Get All Tasks${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/todo"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/todo\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}2.3 Get Task by ID${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/todo/:id"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/todo/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}2.4 Update Task${NC}"
echo -e "${YELLOW}Endpoint:${NC} PUT /api/v1/todo/:id"
echo "curl -X PUT \"${BASE_URL}${API_VERSION}/todo/1\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"title\": \"Updated Task Title\","
echo "    \"description\": \"Updated description\","
echo "    \"status_id\": 2"
echo "  }'"
echo ""

echo -e "${YELLOW}2.5 Delete Task${NC}"
echo -e "${YELLOW}Endpoint:${NC} DELETE /api/v1/todo/:id"
echo "curl -X DELETE \"${BASE_URL}${API_VERSION}/todo/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}2.6 Get Tasks by Responsible User${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/todo/responsible/:userID"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/todo/responsible/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}2.7 Get Tasks by Author${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/todo/author/:userID"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/todo/author/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}2.8 Get Overdue Tasks${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/todo/overdue"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/todo/overdue\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""
echo "---"
echo ""

# ============================================================================
# TASK STATUS ENDPOINTS
# ============================================================================
echo -e "${GREEN}3. TASK STATUS ENDPOINTS${NC}"
echo ""

echo -e "${YELLOW}3.1 Create Task Status${NC}"
echo -e "${YELLOW}Endpoint:${NC} POST /api/v1/statuses"
echo "curl -X POST \"${BASE_URL}${API_VERSION}/statuses\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"name\": \"In Progress\","
echo "    \"description\": \"Task is currently in progress\""
echo "  }'"
echo ""

echo -e "${YELLOW}3.2 Get All Task Statuses${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/statuses"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/statuses\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}3.3 Get Task Status by ID${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/statuses/:id"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/statuses/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}3.4 Update Task Status${NC}"
echo -e "${YELLOW}Endpoint:${NC} PUT /api/v1/statuses/:id"
echo "curl -X PUT \"${BASE_URL}${API_VERSION}/statuses/1\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"name\": \"Completed\","
echo "    \"description\": \"Task has been completed\""
echo "  }'"
echo ""

echo -e "${YELLOW}3.5 Delete Task Status${NC}"
echo -e "${YELLOW}Endpoint:${NC} DELETE /api/v1/statuses/:id"
echo "curl -X DELETE \"${BASE_URL}${API_VERSION}/statuses/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""
echo "---"
echo ""

# ============================================================================
# TASK TYPE ENDPOINTS
# ============================================================================
echo -e "${GREEN}4. TASK TYPE ENDPOINTS${NC}"
echo ""

echo -e "${YELLOW}4.1 Create Task Type${NC}"
echo -e "${YELLOW}Endpoint:${NC} POST /api/v1/task-type"
echo "curl -X POST \"${BASE_URL}${API_VERSION}/task-type\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"name\": \"Bug\","
echo "    \"description\": \"Bug fix task\""
echo "  }'"
echo ""

echo -e "${YELLOW}4.2 Get All Task Types${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/task-type"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/task-type\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}4.3 Get Task Type by ID${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/task-type/:id"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/task-type/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}4.4 Update Task Type${NC}"
echo -e "${YELLOW}Endpoint:${NC} PUT /api/v1/task-type/:id"
echo "curl -X PUT \"${BASE_URL}${API_VERSION}/task-type/1\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"name\": \"Feature Request\","
echo "    \"description\": \"New feature request\""
echo "  }'"
echo ""

echo -e "${YELLOW}4.5 Delete Task Type${NC}"
echo -e "${YELLOW}Endpoint:${NC} DELETE /api/v1/task-type/:id"
echo "curl -X DELETE \"${BASE_URL}${API_VERSION}/task-type/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""
echo "---"
echo ""

# ============================================================================
# WORKFLOW ENDPOINTS
# ============================================================================
echo -e "${GREEN}5. WORKFLOW ENDPOINTS${NC}"
echo ""

echo -e "${YELLOW}5.1 Create Workflow${NC}"
echo -e "${YELLOW}Endpoint:${NC} POST /api/v1/workflows"
echo "curl -X POST \"${BASE_URL}${API_VERSION}/workflows\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"name\": \"Development Workflow\","
echo "    \"description\": \"Standard development workflow\""
echo "  }'"
echo ""

echo -e "${YELLOW}5.2 Get All Workflows${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/workflows"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/workflows\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}5.3 Get Workflow by ID${NC}"
echo -e "${YELLOW}Endpoint:${NC} GET /api/v1/workflows/:id"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/workflows/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""

echo -e "${YELLOW}5.4 Update Workflow${NC}"
echo -e "${YELLOW}Endpoint:${NC} PUT /api/v1/workflows/:id"
echo "curl -X PUT \"${BASE_URL}${API_VERSION}/workflows/1\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{'"
echo "    \"name\": \"Updated Workflow\","
echo "    \"description\": \"Updated workflow description\""
echo "  }'"
echo ""

echo -e "${YELLOW}5.5 Delete Workflow${NC}"
echo -e "${YELLOW}Endpoint:${NC} DELETE /api/v1/workflows/:id"
echo "curl -X DELETE \"${BASE_URL}${API_VERSION}/workflows/1\" \\"
echo "  -H \"Content-Type: application/json\""
echo ""
echo "---"
echo ""

# ============================================================================
# SWAGGER DOCUMENTATION
# ============================================================================
echo -e "${GREEN}6. SWAGGER DOCUMENTATION${NC}"
echo ""
echo -e "${YELLOW}View Swagger UI:${NC}"
echo "curl -X GET \"${BASE_URL}/swagger/index.html\""
echo ""
echo -e "${YELLOW}Get Swagger JSON:${NC}"
echo "curl -X GET \"${BASE_URL}/swagger/doc.json\""
echo ""
echo "---"
echo ""

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}         NOTES AND USAGE TIPS          ${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo "1. Replace IDs (1, 2, etc.) with actual IDs from your database"
echo "2. Modify JSON payloads according to your actual data model"
echo "3. Add headers as needed (e.g., Authorization tokens)"
echo "4. Use -v flag for verbose output: curl -v ..."
echo "5. Use -i flag to include headers in the response: curl -i ..."
echo "6. Use -d with -s for silent mode: curl -s -d '...' ..."
echo "7. Use jq for JSON formatting: curl ... | jq"
echo ""
echo "Example with jq formatting:"
echo "curl -s -X GET \"${BASE_URL}${API_VERSION}/todo\" | jq"
echo ""
echo "Example saving response to file:"
echo "curl -X GET \"${BASE_URL}${API_VERSION}/todo\" -o response.json"
echo ""
echo -e "${BLUE}========================================${NC}"
