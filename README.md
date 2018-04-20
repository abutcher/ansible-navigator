# Ansible navigator
  Command-line utility to navigate ansible playbooks and roles

  Usage:
  ```
  ansible-navigator FILENAME ROW COL
  ```

  Result:
  - blank if position does not correspond to a linked file or said file cannot be found
  OR
  - full path to file that is being referenced
