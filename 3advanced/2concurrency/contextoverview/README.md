# Context Overview

## What is Context?
- Context is a package that helps with communication between processes

## When to use Context?
- General data attachment
  - Auth
  - Request ID
  - Don’t overuse
- Timeout
  - When a context should end
  - Is a context active/cancelled/past deadline
    - You can proceed with the process with different sets of behaviors base on the status
    
## How to use Context?
- ctx context.Context as first if they are multiple parameters
- Make sure the attached context is properly wrapped up (cancelled) before existing the process (process may end long before the originally set timeout)