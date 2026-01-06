---
title: Project Vaxitas: 2026
summary: Listen to me yap about my motivation for maintaining this site, why my site supports more features than I've used, how I achieved a breakthrough once I created the snippets section, and my vision for the site in 2026.
categories: personal, indieweb
date: 2026-01-06T16:43:00+05:30
---

When I first shipped Vaxitas last July, I had published a [spec](about/vaxitas) that entailed the purpose the site served for me and my audience, which I envisioned to be primarily a bunch of friends and family whose company I enjoy, along with a few others who are invested in my life. My overarching goals and intentions with the site remain the same. I want it to be a space on the web where I spend a significant amount of my time reflecting, unabashedly expressing ideas, and morphing them into newer forms in the presence of previously unavailable information.

For my site to provide a conducive environment in which I can play around with ideas, I needed the interface to reflect and support the ways in which I wanted to think and write. After a lot of deliberation, both the site and the site generator that powers it (launching soon™) were constructed with these requirements in mind. But, alas, I'm now in a predicament in which my site offers more features than I've actually utilized in the past year. 

Although I didn't write as much as I would have liked to, I did hit publish a few times. In my defense, the sheer quantity of posts isn't as important to me as taxonomy is. I wanted the ability to structure my posts as individual ideas and clearly establish an individual post's relationships with other relevant ones so that it's easier to navigate around the site. And this feature was severely underutilized because I overwhelmed myself into inaction. 

One of my biggest motivations for maintaining this site is my belief that it can double as a serendipity engine. I have a knack for making a weak first impression both in online spaces and in-person group settings. I like to hang back and listen (usually in adoration when it's my friends) to what the others have to say. I too have plenty of things that I want to talk about. It's just that I think it might not be everybody's cup of tea. 

But, I was certain that there would at least be a handful of people who might share similar interests, and that was enough encouragement for setting up a space in which I could write about them without fear of judgement. Sadly, I would constantly get paralyzed whenever I decided to write. I felt that my ideas lacked the level of polish, peculiarity, or originality to exist as an essay on my site.

After a few months, I tackled this paralysis by creating a separate space on my site — [snippets](/snippets). I'm still unsure about what kind of posts this space encompasses. Initially, I thought I would house posts that I could write in a single sitting, unoriginal ideas that excite me but countless others have already spoken about, and notes on things that I'm learning. Currently, I just use it as my playground. I like experimenting and thinking about how I'd like to structure my ideas and group them, whether I should revise certain posts as I learn more about the subject matter, etc. I had my highest throughput (and viewership) in Dec '25. I wrote 4 snippets that month. And it didn't feel energy-sapping. I was actually having fun. I decided to write about low-stakes stuff so that I didn't spend weeks on a single post (I'm a notoiously slow writer, both in terms of prose and code), and that really helped! I experimented with my voice, the layouts, and it's led to more ideas about what I want to do with the space and what I want to reserve the blog section for. I want to use it for more original pieces of writing. Even if the essence of the idea itself is trite, if I can make it ancedotal or personalized enough, I consider that to be an original.

I've built a decent base for my site over the past year, and I'm excited for how it will shape up in 2026. There are some new things that I'd like to try, and some things I'd like to build on further. 

- **Snippets will be frequently revised**, and a list of the most recently edited ones will be provided at [/snippets](/snippets). 
  - Ex. I want to write about my experiences with setting up and working in monorepos and I see myself revisiting this snippet as my understanding of this topic deepens.
- I want to structure some of my posts as series. For each series, there can exist a home page that describes what it's about and then lists out the main canon posts in sequence. These series could also have filler posts that may not have been written with the intention of fitting into the canon, but are important to the lore.
  - Ex. I want to have an SICP series, in which I talk about my main learnings after each chapter. But, there are certain problems such as the [square root one](/breaking-square-root) which I'd written about and the "coin change" one. These are fun problems that required some thinking, and I feel that they deserve a dedicated write-up. Although they don't fit into the overall theme of the series, they're still relevant to my experience with the book.
- I want to **write interactive posts**. The closest example I possess in my corpus is the [snippet on masonry layouts](/masonry-with-css-and-js). But not quite — the iterations of the component throughout the snippet are mostly static.
  - I have an idea for a series on interesting UI components and how one could make them from scratch. I want to use [Emil](https://emilkowal.ski/ui/the-magic-of-clip-path) and [Nanda](https://www.nan.fyi/database)'s posts as references. 
  - The best way to achieve this is by using custom components in `.mdx` files instead of `.md` and configure [`mdsvex`](https://mdsvex.pngwn.io/), a tool that compiles Markdown (along with embedded Svelte components) down to HTML, to detect these custom Svelte components, and actually insert them in my pages. I'll need to grapple with this problem to figure out how I can enable this in my existing set-up.
- I want to **maintain lists**. Lists of all sorts: food I like, games I've played, media I've consumed, activities and hobbies I've picked up, etc. I'll create a separate space for them.
- **Maintain a list of recommended posts and popular posts in a section on the home page**, so that I can direct new readers to writing that they're likelier to find interesting. This seems like a nice way to welcome them to my site. I'll get on this in a couple of months once I have a sizable corpus.
- **Maintain separate RSS feeds for each section on my site**. The main feed will only consist of the 10 or 15 most recent posts across all of these sections.

This year, I don't foresee myself writing on such a grand scale that I would need to create a sophisticated filter system for my posts. We'll deal with it when we get there.

*This essay was written from a prompt in the [IndieWebClub Meetup, Bangalore](https://underline.center/c/indiewebclub/10).*

> What themes and/or intentions do you have around your blog/personal website in 2026?
