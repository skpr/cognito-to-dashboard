Cognito to CloudWatch Dashboard
===============================

## Goal of this Project

A simple implementation to streamline access to CloudWatch Dashboards.

## Flow

The following diagram outlines the flow of requests.

```mermaid
sequenceDiagram
    Browser->>+API: Browse to link provided by documenation or CLI.
    API->>+Config: Check dashboard is in allowed list. Store dashboard name keyed by state value.
    API->>-Browser: Return a redirect to the Cognito Hosted UI containing state key.
    Browser->>+Cognito: Login using Hosted UI.
    Cognito->>-Browser: Returns redirect to callback URL with state key.
    Browser->>+API: Send request to callback URL. Validate state key and lookup dashboard name from storage.
    API->>+Cognito: Requests temporary credentials using code converted to token.
    API->>-Browser: Return direct federated login link to CloudWatch Dashboard.
    Browser->>+CloudWatch: Follow federated login link to CloudWatch Dashboard.
```

## Configuration

See [defaults.env](/defaults.env) for environment variables.
