---
slug: "/learn/agentic-coding"
nav: "the loop"
blurb: "vibe coding, the Ralph loop, agentic engineering: how this very site builds itself."
teaches: "agentic coding, the Ralph loop, vibe coding, agentic engineering"
order: 10
---

The narrator of this school got replaced by an AI agent. Not metaphorically. The pipeline that writes, sanitizes, commits, and deploys every pelican on this site runs inside a loop that re-feeds the same prompt file to a coding agent overnight while the narrator sleeps. You are reading content orchestrated by the very thing being described. Either the most educational conflict of interest in the history of adult learning, or just extremely funny. Welcome to the meta-lesson.

:::beakman
sit DOWN. a coding agent does not have to be RIGHT the first time. *burp* it writes, runs the code, reads the error, fixes itself. that self-debug loop is the whole job.
:::

::ralphloop

::art{id="agentic-coding-2" caption="vibe coding: the bird shrugs, ships it, and refuses to read its own code"}

## vibe coding (2025): fully give in, forget the code even exists

In February 2025, Karpathy posted a tweet that launched a thousand hot takes. He called the new practice **vibe coding**: describe what you want, the AI writes the code, you do not read it, you just run it and see if the vibes are right. His exact phrase: "fully give in to the vibes, embrace exponentials, and forget that the code even exists." Partly a joke, partly a genuine observation that for prototypes and throwaway scripts you really could stop reading your own codebase.

The hype cycle did what hype cycles do: every startup declared programming was over and engineers were obsolete. The engineers kept their jobs and started using the tools to write code faster.

## agentic engineering (2026): the grown-up version

By 2026, Karpathy updated the framing. "Vibe coding" was the gateway drug. The mature practice is **agentic engineering**: you are not writing code 99% of the time. You are orchestrating agents, reviewing output, acting as oversight. Set direction, evaluate results, catch mistakes, decide when to push the button.

The skill is no longer "can you write Python." It is "can you decompose a problem clearly enough that an agent can execute it, and can you tell when it has gone wrong." Somewhat inconveniently for the people who declared engineers obsolete: a higher-order skill, not a lower one.

:::beakman{compact variant="dr-beakman-2"}
"the AI took the jobs"? no. the job moved UP, your little meat-brain: from typing code to deciding when the agent is wrong.
:::

## the autonomy slider

Karpathy describes this as the **autonomy slider** (Software 3.0, YC 2025). At one end, the agent asks about every decision. At the other, it runs for hours without checking in. Neither extreme is right for every job:

- **Low autonomy:** "write me a function that does X, show me the code, I will paste it in." You stay in control. The agent is fast autocomplete.
- **Medium autonomy:** "refactor this module and run the tests; ask me if you hit something ambiguous." You review diffs. The agent does the work.
- **High autonomy:** "here is PROMPT.md and AGENTS.md; build until the tests pass; push when done." You check git in the morning. The agent ran all night.

The dial also controls how often the agent hits its context limit and starts to degrade, which brings us to the loop.

::art{id="agentic-coding" caption="a pelican agent sprinting a loop, chasing something that stays just out of reach"}

## the Ralph loop: a while-loop as architecture

Geoffrey Huntley figured out something that sounds absurd and turns out to be load-bearing. The Ralph loop is a Bash `while true` that wakes up a coding agent, hands it a `PROMPT.md`, waits for it to finish, and wakes it up again. Forever. Overnight. While you sleep. Ralph Wiggum: a bit simple, a bit earnest, just keeps running.

The clever part is what the loop *solves.* Quality degrades past roughly 100,000 to 150,000 tokens, the "Dumb Zone" where the model is too distracted to reason clearly. Huntley calls a long-running agent that never resets "deterministically bad in an undeterministic world." The loop fixes it: kill the agent, start a fresh context, feed the same spec file. The **filesystem is the memory.** The agent does not need to remember the previous run because all the code it wrote is right there on disk. Fresh context, durable state.

Huntley runs the loop 12 hours overnight. By morning, dozens of incremental commits, each a short coherent run, the codebase moved forward without anyone at the keyboard. Anthropic baked this directly into Claude Code as the built-in `/loop` command.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
the loop is not a hack, it is the fix. the model rots past ~150k tokens, so you kill it and reload the spec. the DISK is the memory, not the context.
:::

## Gas Town: when loops beget infrastructure

One loop begets a flock. Steve Yegge (Amazon, Google, Sourcegraph) spent late 2025 building **Gas Town** (launched January 2026): an open-source system coordinating 20 to 30 Claude Code instances on the same codebase at once. "Kubernetes for AI coding agents," which is architecturally accurate, and roughly $100 an hour to run. The pelican just needs to know this level exists.

## the open-source local toolbox

You do not need Claude Code or Cursor. All of these tools in 2026 can point at a local model via [Ollama or LM Studio](/pelicanplus/local) so no tokens leave your machine:

- **OpenCode** (the most-starred Claude Code alternative in 2026): terminal-native, model-agnostic.
- **Cline**: a VS Code extension with a full autonomous agent mode and a community of power users.
- **OpenHands** (formerly Devin-open): a sandboxed autonomous agent that can browse, run code, and commit.
- **Aider**: git-native pair programmer; every change is a diff you can review before committing.
- **Goose**: Block's open-source autonomous coding agent (Apache 2.0, now governed by the Linux Foundation's Agentic AI Foundation); works with any LLM provider including local models via Ollama.
- **Codex CLI**: OpenAI's terminal agent, open-sourced in 2025.

All of them accept a `PROMPT.md` or equivalent spec file, and any of them can be the thing inside the Ralph loop. Whether you point them at a rented frontier bird or one you own outright is the next lesson's whole argument.

## the meta payoff: this site is a Ralph loop

Pelicans.wtf has a load-bearing `PROMPT.md` specifying exactly how to generate, sanitize, and describe a pelican SVG, and an `AGENTS.md` documenting the codebase for any agent working in the repo. When a new model drops, the curator runs `npm run generate-next` and walks away. The pipeline calls the model, sanitizes the output, writes a description, commits, pushes. The push is the deploy. Nobody typed the commit message. Nobody reviewed the SVG before it went out. The agent ran the loop.

The narrator got replaced by an agent, started a website about AI, and is now running an agent to build the website. The irony is the entire point. This school exists because the people most qualified to explain agentic coding are the ones who got automated out of a job and had to use the same tools to build something new. The pelican on the bicycle is not just a benchmark. It is also a mood.

::art{id="agentic-coding-4" caption="Gas Town: a flock of coding-agent pelicans swarming one codebase at once"}

:::beakman{variant="dr-beakman-4"}
a text-predictor that learned to debug itself, on a fresh context, at 3am, while you drool on your pillow. fine, it is genuinely amazing. do not repeat that.
:::

:::sources
**sources:** the loop keeps running. the sources are real.

- Andrej Karpathy, ["Software Is Changing (Again)"](https://www.youtube.com/watch?v=LCEmiRjPEtQ) (YC AI Startup School, 2025): Software 3.0, the autonomy slider, vibe coding to agentic engineering.
- Geoffrey Huntley, [ghuntley.com/ralph/](https://ghuntley.com/ralph/) (the primary source for the Ralph loop: while-true architecture, filesystem as memory).
- Yao et al. (2022), ["ReAct: Synergizing Reasoning and Acting in Language Models"](https://arxiv.org/abs/2210.03629) (the reason-then-act loop underneath every coding agent in this lesson).
- Jimenez et al., ["SWE-bench: Can Language Models Resolve Real-World GitHub Issues?"](https://arxiv.org/abs/2310.06770) (arXiv, 2023): the benchmark that measures agentic coding against actual GitHub issues.
- OpenHands, [docs.openhands.dev](https://docs.openhands.dev) (open-source sandboxed agent that browses, runs code, and commits; works with any LLM).
:::
