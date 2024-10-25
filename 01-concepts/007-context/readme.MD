Why context?

- It simplifies the implementation for cancellation across API and processes boundaries.
- To signal cancellation.
- Prepares your code base for scale.
  - how to I cancel it efficiently.

- it carries deadlines, cancelation signals and other request scoped values across API boundaries and between processes.

Functions
- WithCancel
- WithDeadline
- WithTimeout
- WithValue

