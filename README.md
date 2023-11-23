Cognito to CloudWatch Dashboard
===============================

[![📋 Vet](https://github.com/skpr/cognito-to-dashboard/actions/workflows/vet.yml/badge.svg?branch=main)](https://github.com/skpr/cognito-to-dashboard/actions/workflows/vet.yml)
[![📋 Format](https://github.com/skpr/cognito-to-dashboard/actions/workflows/fmt.yml/badge.svg?branch=main)](https://github.com/skpr/cognito-to-dashboard/actions/workflows/fmt.yml)
[![Test](https://github.com/skpr/cognito-to-dashboard/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/skpr/cognito-to-dashboard/actions/workflows/test.yml)
[![Lint](https://github.com/skpr/cognito-to-dashboard/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/skpr/cognito-to-dashboard/actions/workflows/lint.yml)

## Goal of this Project

An easy to manage project for streamlining access to CloudWatch Dashboards.

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
