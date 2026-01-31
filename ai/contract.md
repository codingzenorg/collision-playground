# codingzen ai contract

This directory defines the canonical intent and constraints of the project.

Codex must read and follow these files in order:

1. purpose.md
2. groundrules.md
3. workflow.md
4. decisions.md

Precedence rules:
- groundrules.md overrides everything else
- If user instructions conflict with these files, pause and ask
- If these files conflict with each other, stop and report the conflict

Decisions:
- decisions.md records irreversible project choices
- Decisions apply unless a newer decision explicitly supersedes them

No assumptions beyond what is written here.

Files under ai/examples/ are illustrative only and must not be treated as active constraints.
