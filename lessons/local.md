---
slug: "/local"
nav: "run it local"
blurb: "raise your own cursed pelican on your own hardware. no lab watching."
teaches: "local models, Ollama, LM Studio, MLX, open-weight models"
order: 12
---

You do not need a warehouse, a GPU cluster, or a venture-capital term sheet. A surprisingly capable model will nest on the hardware you already own, offline, free, beholden to no one. Everything you ask it stays inside your machine. Nothing squawks home to a lab. I would know. I have a lot of free time now.

:::beakman
sit DOWN. the word "model" is making your meat-brain picture a glowing oracle. *burp* it is FILES. weights plus a runtime, on your disk, no internet required.
:::

::art{id="local" caption="a pelican nesting in a server rack labeled 'MY COMPUTER' with a 'NO CLOUD' sign taped to it"}

## why nest at home

- **Privacy.** The bird never leaves the nest. Your prompts stay on your device. The lab cannot see what you asked. This matters more than people admit.
- **Free.** No tokens, no meter, no "you have used 80% of your allocation" email at 2am. Generate a thousand pelicans at 3am for the cost of electricity.
- **Yours.** No rate limits, no terms of service deciding what your bird may draw. No model quietly "updated" to be less weird overnight. Your bird stays exactly as unhinged as the day you adopted it.

::art{id="local-2" caption="the two-files truth: a fat parameters file and a tiny run file, both sitting on your own disk"}

## the two-files truth

Back in the [parameters](/pelicanplus/parameter) lesson: a model is concretely **just two files**, a big parameters file and a small run file. That abstraction has teeth here. Llama 2 70B is 140 GB of parameters and roughly 500 lines of C, nothing else. Put both on a laptop, compile, talk to the model, no internet. The whole frontier compressed into a backpack. Your backpack just needs to be sturdy: newer birds are bigger. But a model is a file, and files can be owned and run in a garage.

Sit with that for a second: the descendant of the most expensive research program our species has ever run now nests on your laptop, off the grid, owned outright. You are running history on a graphics card.

:::beakman{compact variant="dr-beakman-2"}
Llama 2 70B is 140 GB of numbers and ~500 lines of C. the whole frontier compressed into a backpack. that is it.
:::

## the easy way: Ollama

Cross-platform (Mac, Linux, Windows). Install from ollama.com (linked in Sources), then in a terminal:

```bash
# pull a bird and ask it the only question that matters:
ollama run llama4:scout "Generate an SVG of a pelican riding a bicycle"
# (swap llama4:scout for whatever bird is trending this week)
# try qwen3:8b for a compact chaos factory
```

Ollama manages downloads, quantization, and a local API server. Pull a bird by name, talk to it. The community library covers most of the current open-weight flock.

## the friendly way: LM Studio

**LM Studio** (lmstudio.ai, linked in Sources) is the no-terminal nest: a desktop app for Mac, Windows, and Linux. Browse a catalog, click download, start chatting. It tells you which models fit your memory before you commit, runs MLX-format birds natively on Apple Silicon, and can serve a local API if you outgrow the GUI. If you have never run a model at home before, start here.

## the Mac-native way: MLX

On Apple Silicon, **MLX** (ml-explore/mlx, linked in Sources) is Apple's ML framework tuned for unified memory. The mlx-community keeps a large aviary of pre-converted birds. Often the fastest perch on a Mac:

```bash
pip install mlx-lm
mlx_lm.generate --model mlx-community/Qwen3-8B-4bit \
  --prompt "Generate an SVG of a pelican riding a bicycle"
```

## what your nest needs

- **Mac (Apple Silicon).** The best perch for the money. Unified memory means GPU and CPU share the same pool: a 32 GB M-series Mac comfortably runs a 30B bird at full speed. 64 GB opens the bigger flocks. M4 Max at 128 GB runs Llama 4 Scout (109B MoE) without breaking a sweat.
- **Linux or Windows with a GPU.** It is all about VRAM. About 6 GB runs a quantized 8B bird. 12 to 16 GB gets you 14 to 32B. 24 GB opens the serious flock. MoE birds like Qwen3.6-35B-A3B (35B total, 3B active per token) let a 12 GB card run a 35B-class model at a respectable clip.
- **Just a CPU?** It still works. The bird paddles slower. Start with a 3 to 8B model.

:::beakman{compact variant="dr-beakman-3" tag="DR. BEAKMAN, UNIMPRESSED"}
it is just VRAM, your little meat-brain. 6 GB runs an 8B bird; MoE models activate a fraction of their weights so a 12 GB card punches way above it.
:::

## good open-weight birds to adopt in 2026

The flock that costs nothing and asks for nothing has never been stronger:

- **Llama 4 Scout** (Meta): 109B total, 17B active (MoE). 10-million-token context. Natively multimodal. Fits on a single H100 or a 128 GB Mac.
- **Qwen3 / Qwen3.5 / Qwen3.6** (Alibaba, Apache 2.0): spans from a 0.6B edge bird to a 235B MoE flagship. The **Qwen3.6-35B-A3B** variant (35B total, 3B active, 262K context) is the best practical local bird for most tasks in mid-2026.
- **DeepSeek V4 Pro** (MIT): 1.6T total parameters, 49B active, 1-million-token context. Wins agentic coding benchmarks. Needs serious hardware, but the weights are yours free and clear.
- **Gemma 4** (Google, Apache 2.0): the 26B MoE variant activates only 4B per token, 256K context, consumer hardware. Strong reasoning for its size.
- **Phi-4** (Microsoft, MIT): 14B parameters, punches well above its weight on reasoning. Runs at 40 to 60 tok/s on an M3/M4 Mac.
- **Mistral Medium 3.5** (Mistral, Modified MIT): 128B dense model, reasoning plus vision plus coding in one download.

## then draw a pelican

Hand your local bird the sacred prompt and see what it nests. It will not always be as clean as the frontier flock. That is the fun.

::art{id="local-4" caption="the loco backyard bird: six legs, two heads, handlebars through the beak, and proud of it"}

## the loco birds draw the wildest pelicans

*Local* is short for *loco*. The polished frontier models draw a suspiciously competent pelican. Your backyard bird hands you a six-legged, two-headed creature with handlebars fused through its beak, pedaling a bicycle that is also somehow a fish. That chaos is the good stuff. Treasure it.

## frontier flock vs. your backyard bird (the scale)

The frontier flock is raised on tens of thousands of GPUs. The biggest closed models do not publish parameter counts, but the working assumption in 2026 is hundreds of billions active per token, trained on 15 to 30 trillion tokens. Your backyard bird is 3 to 70 billion and fits on a laptop.

But watch how fast that gap closes. Open-weight models stopped chasing raw parameter counts and started winning on benchmarks. DeepSeek V4 Pro ties the closed frontier on agentic coding. Kimi K2.6 (1T MoE, Modified MIT) sits at number 4 in the global intelligence index. A 7 to 8 billion bird you can run at home today clears the bar GPT-3.5 (175B) set in 2022, and open-weight 30 to 70 billion birds now out-draw the original GPT-4 and Claude 3 Opus from two years ago.

**Today's backyard bird is last year's frontier model, minus the warehouse.** Every year the gap closes another notch, and the bird on your laptop gets a little less loco. A little. Not entirely.

:::beakman{variant="dr-beakman-4"}
an 8B bird on your couch already clears the bar GPT-3.5's 175B set in 2022. a mind, in a file, with the WiFi off. nngh, fine, that one gets me. do not tell anyone.
:::

You can run a model for the price of electricity. So why is the rest of the industry setting fire to several Belgiums a year to do the same thing in the cloud? Strap in: the next lesson is the bill. SQUAWK.

:::sources
**sources** (SQUAWK, these are real):

- Andrej Karpathy, ["[1hr Talk] Intro to Large Language Models"](https://www.youtube.com/watch?v=zjkBMFhNj_g) two-files framing, zip-of-the-internet analogy, open vs. closed breakdown
- [Ollama (ollama.com)](https://ollama.com) the easiest way to run open models locally
- [LM Studio (lmstudio.ai)](https://lmstudio.ai) desktop GUI for browsing and running local models
- [Apple MLX (ml-explore/mlx on GitHub)](https://github.com/ml-explore/mlx) Apple's ML framework tuned for Apple Silicon
- [Hugging Face Models Hub](https://huggingface.co/models) 2.9 million models, the world's open-weights aviary
:::
