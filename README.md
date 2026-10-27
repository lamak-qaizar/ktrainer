# ktrainer

Strict Red-Green-Refactor TDD kata coach. You specify every test, implementation and refactor; the agent only types what you said and runs the tests.

## Install the skill (Claude Code / Codex / opencode)

```sh
npx skills add lamak-qaizar/ktrainer
```

Pick your agent when prompted, then run `/ktrainer` (or ask for a TDD kata) inside your kata project.

<details>
<summary>No Node? Manual install</summary>

```sh
# Claude Code
mkdir -p ~/.claude/skills/ktrainer && curl -fsSL https://raw.githubusercontent.com/lamak-qaizar/ktrainer/main/skills/ktrainer/SKILL.md -o ~/.claude/skills/ktrainer/SKILL.md

# Codex / opencode
mkdir -p ~/.agents/skills/ktrainer && curl -fsSL https://raw.githubusercontent.com/lamak-qaizar/ktrainer/main/skills/ktrainer/SKILL.md -o ~/.agents/skills/ktrainer/SKILL.md
```

Restart the agent after installing.
</details>

## Install the CLI via scoop

```sh
scoop install https://github.com/lamak-qaizar/ktrainer/releases/latest/download/ktrainer.json
```
