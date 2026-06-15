---
slug: "/learn/context-window"
nav: "context window"
blurb: "the bird's tiny working memory. everything it can hold in its head at once."
teaches: "context window, tokens, attention, working memory"
order: 5
---

The model has two kinds of memory. People mix them up constantly, then get annoyed at the bird for "forgetting," then paste the same PDF in a fifth time. The **parameters** are billions of dial positions baked in during training: blurry long-term recall. The bird knows roughly what a pelican is; it cannot quote the exact sentence it read in 2021. The **context window** is the live, running sequence of tokens the model can see *right now*, fed into the network with zero fuzziness. Your message is in there. The whole chat history is in there. This is the bird reading off a scroll that keeps growing until the chat ends. Then the scroll disappears. SQUAWK.

:::beakman
sit DOWN. parameters are what the bird memorized; the context window is what it can literally SEE right now. confuse them and you will keep yelling at a bird for forgetting things you never put on the scroll.
:::

::art{id="context-window" caption="a pelican with an enormous pouch stuffed with fish-shaped tokens, looking strained"}

::art{id="context-window-2" caption="two memories, two birds: frozen training weights versus the live scroll"}

## the pouch only holds so many fish

The context window has a hard ceiling in tokens. Early models (GPT-2 era) maxed out at 1,024. GPT-4 shipped at 8,192, then 32,768. By 2026 the arms race has produced genuinely absurd pouches: Llama 4 Scout supports **10 million tokens**, and Claude Sonnet 4.6 and Opus 4.8 expanded to **1 million tokens**, generally available since early 2026. Whether you can usefully fill a 10-million-token pouch is a different question. The bird is not necessarily paying full attention to every fish at the back.

:::beakman{compact variant="dr-beakman-2"}
a 10-million-token pouch does not mean it READS 10 million tokens. stuff buried in the middle gets ignored. lost in the middle, your little meat-brain. it is real.
:::

## why this matters for trusting the bird

Same split, sharper: parameters are something you read a month ago (gist, maybe a wrong detail, no exact quote); the context window is the document open on your desk right now, every line readable. So when a model uses web search, it is not "going smarter," it is pulling real text onto the desk so it can read instead of recall. A librarian handing the bird a printout. Hold that thought: it is the entire fix in two later lessons.

::art{id="context-window-3" caption="a fish lost in the middle of a very long scroll"}

## new chat wipes the pouch

Click "new chat" and the context window resets to zero. The parameters survive, permanent, untouched. The conversation does not. Gone. Every preference you established, every file you pasted, the entire backstory you spent forty minutes explaining: gone. The pouch was physically emptied. This is why long-running projects need you to re-introduce context each session, and why the AI's advice in session 1 and session 2 can differ: same bird, empty pouch, slightly different fishing trip.

## irrelevant tokens are a tax

More tokens costs you two ways. Literally: most APIs charge per token. Subtly: irrelevant tokens **distract** the model and lower accuracy. The **attention mechanism** (the core idea Ashish Vaswani and his coauthors introduced in 2017's "Attention Is All You Need," the paper every modern model is built on) looks across everything in the window at once; filling it with noise is like asking someone to find a key fact buried in a pile of unrelated meeting notes. Treat the context window as a precious resource. Keep it short. Keep it on-topic. Start a fresh chat when you switch subjects. Your wallet and your accuracy will both thank you. SQUAWK.

::art{id="context-window-4" caption="the new-chat button emptying a pelican's pouch back to zero"}

:::beakman{variant="dr-beakman-3"}
every token attends to every other token, so cost grows with the SQUARE of the length: double the window, QUADRUPLE the work. *burp* that is why big context is slow and pricey. it is geometry, not laziness. gorgeous geometry, do not tell anyone.
:::

:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["How I use LLMs"](https://www.youtube.com/watch?v=EWvNQjAaOHw) (working-memory framing, when to start a new chat, context curation)
- Vaswani et al. (2017), ["Attention Is All You Need"](https://arxiv.org/abs/1706.03762) (the Transformer paper that defines the attention mechanism underlying context windows)
- Liu et al. (2023), ["Lost in the Middle: How Language Models Use Long Contexts"](https://arxiv.org/abs/2307.03172) (why burying facts in the middle of a long context degrades accuracy)
- Anthropic, ["Context windows" (official API docs)](https://docs.anthropic.com/en/docs/build-with-claude/context-windows) (token limits, counting API, and model-specific context window sizes)
:::
