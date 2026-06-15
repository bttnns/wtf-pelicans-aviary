---
slug: "/learn/token"
nav: "tokens"
blurb: "the atom of everything. the bird does not read words, it reads tokens."
teaches: "tokens, tokenization, byte pair encoding"
order: 1
---

Here is the first thing nobody tells you, and it is load-bearing for everything else in this school: **the bird cannot read.** When you type *pelican* into a model, it does not see seven proud letters. It sees a couple of numbered chunks called **tokens**, and it only ever eats and regurgitates tokens. Never a letter. Never a whole word. Always a beakful of token.

:::beakman
sit DOWN. the bird never sees letters, your little meat-brain. it eats numbered chunks and pukes numbered chunks, end of story.
:::

::art{id="token" caption="a pelican eating a word that has been chopped into token-pieces"}

::art{id="token-2" caption="the word 'pelican' on a chopping board, sliced into numbered chunks"}

## a token is a chunk, not a letter

A token is a *piece* of text: sometimes a whole short word, usually a fragment, sometimes punctuation or whitespace. The model has a fixed menu of them, called its **vocabulary**. Every bit of text you send gets minced into menu items before the bird tastes any of it.

GPT-2 (2019) ate from a menu of about **50,000 tokens**. GPT-4 (2023) grew that to **100,000**. GPT-4o (2024) doubled it to roughly **200,000**. More menu, bigger bites: the same sentence becomes fewer, fatter tokens and the bird sees more text before its mouth is full. The pouch keeps getting roomier.

## where the menu comes from (BPE)

Nobody hand-wrote 200,000 tokens. The menu is grown by an algorithm with the gloriously unglamorous name **byte pair encoding**. Rico Sennrich and colleagues brought it into language modeling in 2016 to handle rare words; Alec Radford and the OpenAI team carried the same trick into the GPT line, which is why every model you talk to today still eats from a BPE menu. The recipe:

1. Start with every character as its own tiny token.
2. Find the two neighbors that appear together most often. Fuse them. Add to the menu.
3. Repeat tens of thousands of times.

Common pairs like "th" and "ing" fuse early. Rare combos stay as crumbs. Common English words become single tokens; unusual words get shredded into pieces; emoji become a whole adventure.

:::beakman{compact variant="dr-beakman-2"}
nobody wrote the menu. greediest pairs fuse first, raw statistics carve the rest. *burp*
:::

::art{id="token-3" caption="a giant menu board of token-chunks, common pairs fused together"}

## the strawberry problem (and why it mostly got fixed)

For a few years every AI demo included the trick: ask a model how many r's are in *strawberry*. Early models said two. The reason was tokenization: "strawberry" got chopped into two or three tokens, letters inside a token pureed beyond recognition. You were asking the bird to count sprinkles blended into a smoothie. By 2024-2025 the labs patched this through larger vocabularies and reasoning-focused fine-tuning. Modern models usually get it right. But the lesson stands: a startling amount of model "dumbness" is tokenization having a moment. The blender is still running. It is just a fancier blender now.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
it cannot count r's because it never saw letters, it saw a smoothie. that is not dumb, that is tokenization.
:::

::art{id="token-4" caption="a pelican peering into a blender full of strawberry smoothie, trying to count the seeds"}

## why you should care

Tokens are the unit of *everything* downstream. The model thinks in tokens, its memory is measured in tokens, and your bill, the one that replaced your salary, is counted in tokens. The bird is not reading. It is pattern-matching on a menu it memorized during training, and doing it frighteningly well for something that has never seen the alphabet.

So: the bird eats tokens. Next question, the one the whole school turns on: where does it keep what it learned? The answer is a pile of numbers nobody set by hand. SQUAWK.

:::beakman{variant="dr-beakman-4"}
a thing that eats word-confetti can write you a sonnet. that should not work, and it does, and it is gorgeous, okay? do not tell anyone I said that.
:::

:::sources
**sources, because a bird is not a peer-reviewed citation:**

- Andrej Karpathy, ["Let's build the GPT Tokenizer"](https://www.youtube.com/watch?v=zduSFxRajkE) and ["Deep Dive into LLMs like ChatGPT"](https://www.youtube.com/watch?v=7xTGNNLPyMI) (builds the tokenizer from scratch; source of the strawberry explanation)
- Sennrich, Haddow, Birch (ACL 2016), ["Neural Machine Translation of Rare Words with Subword Units"](https://arxiv.org/abs/1508.07909) (the paper that introduced BPE to NLP)
- OpenAI, [tiktoken (GitHub)](https://github.com/openai/tiktoken) (the fast BPE tokenizer used by GPT models; cl100k for GPT-4, o200k for GPT-4o)
- Hugging Face, [Tokenizers library documentation](https://huggingface.co/docs/tokenizers/en/index) (training and running tokenizers in research and production)
:::
