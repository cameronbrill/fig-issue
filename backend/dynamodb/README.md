# data model

## access patterns

create ticket for user

get tickets for user, grouped by:

- import connection
- outport connection

get connections:

- input:
  - name, owner: user, meta: Figma | (other), created_on
    - Figma:
      - api key, team, uniquePasscode
- output:
  - name, owner, meta: Linear | ShortCut | (other), created_on
    - Linear:
      - team, linear api key
- link:
  - input: input, output: output

get events:

- name, source, created_at, input_source: connection<Figma | (other)>, Figma, output_meta: list[Linear | (other)]
  - Figma
    - type: comment, created_at, resolved_at, comment: list[str], mentions: list[str]
  - Linear
    - type: issue, created_at, id, connection<Linear | (other)>
