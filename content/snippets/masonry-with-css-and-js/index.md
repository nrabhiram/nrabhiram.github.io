---
title: Masonries For Dummies
summary: I was ideating on how I wanted my snippets page to look, and I used a tumblr archive as a reference. I wanted variable heights, which turned out to be a task and a half because there wasn't a native CSS solution, and I learned that this layout had a well-established name: masonry.
preview: <div class="content"><img src="/content/snippets/masonry-with-css-and-js/preview.png" alt="Masonry preview" loading="lazy"><p>I was ideating on how I wanted my snippets page to look, and I used a tumblr archive as a reference. I wanted variable heights, which turned out to be a task and a half because there wasn't a native CSS solution, and I learned that this layout had a well-established name: masonry.</p></div>
categories: programming, web
date: 2025-12-15T01:39:00+05:30
cast-hash: 0xdb69ab36
---

For months, I'd been putting off creating this part of my site. I wanted a space where I could write essays more freely without having to mull over the level of polish. I took inspiration from [Judah's riffs](https://joodaloop.com/riffs/), [Abhinav's notes](https://abhinavsarkar.net/notes/), and [Brandur's fragments](https://brandur.org/fragments). I wanted a name for this space that was original. After some deliberation, it finally came down to two contenders: *quickies* and *snippets*. I decided to go with the latter because it wouldn't be perceived as a sexual innuendo.

## Tumblr has Cool Layouts

I used this [tumblr archive](https://lifeafterpsychiatry.tumblr.com/archive) (my friend had shared) as a reference for how I wanted the [snippets](/snippets) page to look. I liked the idea of a grid, but I didn't want it to look uniform. I wanted to make the individual posts look like post-it notes or polaroids, and I wanted the height of posts in a row to be variable.

The simplest solution for this is to use the `columns` CSS property.

```css
.container {
    columns: 3;
}

.container .item {
    break-inside: avoid;
}
```

<div class="masonry-container with-cols">
    <div class="item" style="height: 60px; margin-top: 0px;">1</div>
    <div class="item" style="height: 60px;">2</div>
    <div class="item" style="height: 60px;">3</div>
    <div class="item" style="height: 60px;">4</div>
    <div class="item" style="height: 60px;">5</div>
    <div class="item" style="height: 60px;">6</div>
    <div class="item" style="height: 60px;">7</div>
</div>

The problem this presents is that the flow of the items is from top to bottom in a column, and then to the next column on the right. But, when you visually process a grid of items, you move from left to right, and then downwards — the exact opposite. This is jarring. I can't expect my readers to scroll up and down to view all of my posts sequentially.

## Masonry

There isn't a native CSS solution for this, but there's a hacky one that works perfectly — provided you're okay with writing a bit of CSS and minimal JS. And it turns out, there's already an established name for this component. It's called *masonry*, i.e. stonework. Not to be confused with freemasonry, that's a cult...I think.

**Note**: Before proceeding, I would like to thank Tobias Ahlin for their [implementation guide](https://tobiasahlin.com/blog/masonry-with-css/). It saved me from wasting a great deal of time racking my head about this. In this post, I'm simply trying to articulate what I learned and verify whether I can visualize why this works.

### Order in the box

There are two main principles that we need to understand about `order` in order to understand the implementation. The `order` CSS property, as the name suggests, allows us to order the items in a flexbox, irrespective of the semantic ordering. The lower the `order` value, the higher the priority, i.e. an item is arranged before the other items with higher `order` values. But, if two items have the same `order` value, the one that semantically precedes the other is arranged first.

### How Should the Items Flow

The `flex-flow` CSS property is a shorthand that combines both the `flex-direction` and `flex-wrap` properties. We apply a value of `column wrap` because we want items to be arranged in a column, and if the height of the items exceeds the flex container's height, the last item wraps and is added to the top of the adjacent column on the right.

**Note**: We need to specify a height for the flex container. Otherwise, the criterion for wrapping is never met. Items wrap only when the height of a column exceeds the container's. For now, we'll hardcode the height. But, I'll share a solution for programmatically setting the height towards the end.

### Default Behaviour

Now comes the fun part. The items are currently ordered sequentially but along the column, which yields the same result as using the `columns` CSS property on the container.

```css
.container {
    display: flex;
    flex-flow: column wrap;
    height: 248px;
    width: 100%;
    /* used instead of align-items when content can wrap into multiple lines */
    align-content: flex-start;
}

.masonry-container.default .item {
    margin: 0.5rem;
    width: calc(33.3% - 1rem); /* 0.5rem on either side for gaps */
}
```

<div class="masonry-container default">
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">1</div>
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">2</div>
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">3</div>
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">4</div>
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">5</div>
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">6</div>
    <div class="item" style="width: calc(33.3% - 1rem); height: 60px;">7</div>
</div>

### Flexbox Colomizer

What we want to do is:

- move the 4th and 7th items to the 1st column, below the 1st item (`3n + 1`)
- move the 2nd and 5th items to the 2nd column (`3n + 2`)
- move the 3rd and 6th items to the 3rd column (`3n`)

This is where our knowledge of how `order` works comes in handy. We notice a pattern emerging in the lists of items that should be aligned in a column. We use the `:nth-of-type()` pseudo-class to match these items (you're allowed to express the position of the item in the form `an + b`).

```css
.container .item:nth-of-type(3n + 1) {
    order: 1;
}

.container .item:nth-of-type(3n + 2) {
    order: 2;
}

.container .item:nth-of-type(3n) {
    order: 3;
}
```

<div class="masonry-container default">
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">1</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">2</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">3</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">4</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">5</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">6</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">7</div>
</div>

### Barriers

Okay, this doesn't seem right. `3` should be at the top of the 3rd column. Instead, it's at the bottom of the 2nd one. This is because there's enough space left in the 2nd column to add this item. The columns we have are imaginary, but we can enforce them with some shrewd CSS manipulations. We want to create barriers that can do this for us. How many such barriers do we need? We have 2 separations[^1] between columns: 

- 1 and 2
- 2 and 3

We add a couple of `<span>`s at the end of the container. Why not `<div>`s, you ask? Because the `:nth-of-type()` pseudo-class matches by element type, and both the `div.item` and `div.barrier` elements are considered to be the same type. By using `<span>`s, we don't have to write additional styles for the ordering, because the same generalized rules can be applied.

```html
<div class="container">
    <div class="item">1</div>
    <div class="item">2</div>
    <div class="item">3</div>
    <div class="item">4</div>
    <div class="item">5</div>
    <div class="item">6</div>
    <div class="item">7</div>
    
    <span class="item barrier"></span>
    <span class="item barrier"></span>
</div>
```

The `order` of the first barrier (satisfies `3n + 1`) is `1`. And the `order` of the second barrier (satisfies `3n + 2`) is `2`. Because of the `order` values, the barriers are added to the end of the first and second columns respectively.

Next, we add the following styles for the barriers. The `flex-basis` property defines how much space the flex item will take up.

```css
.container .item.barrier {
    flex-basis: 100%;
    width: 0px;
    margin: unset;
}
```

By setting the value as `100%`, it takes up the full height of the container. This ensures that even if there's some empty space (the amount of empty space is always less than the full height of the container) at the end of our imaginary columns, it can't be illegally occupied by an item that doesn't belong there. Pretty neat, right?

### The Result

*Voila*, it works!

<div class="masonry-container default">
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">1</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">2</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">3</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">4</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">5</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">6</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">7</div>
    <span class="item ordered barrier"></span>
    <span class="item ordered barrier"></span>
</div>

And look, it even supports variable heights!

<div class="masonry-container default" style="height: 280px;">
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">1</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 120px;">2</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">3</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 90px;">4</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">5</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 30px;">6</div>
    <div class="item ordered" style="width: calc(33.3% - 1rem); height: 60px;">7</div>
    <span class="item ordered barrier"></span>
    <span class="item ordered barrier"></span>
</div>

There's just one caveat — we need to hardcode the height for the container. Finding the right height requires a bit of tweaking around in the developer console. But, if you find this annoying, and are willing to write a few lines of JS, this can be handled *programagically*.

```js
function setMasonryContainerHeight() {
  const masonryContainer = document.querySelector('.container');
  if (!masonryContainer) return;

  const masonryItems = Array.from(masonryContainer.querySelectorAll('div.item'));
  if (masonryItems.length === 0) return;

  const columnHeights = [0, 0, 0];

  masonryItems.forEach((masonryItem, index) => {
    const computedStyle = window.getComputedStyle(masonryItem);
    const height = snippet.offsetHeight;
    const marginTop = parseFloat(computedStyle.marginTop);
    const marginBottom = parseFloat(computedStyle.marginBottom);
    const totalHeight = height + marginTop + marginBottom;

    const childNumber = index + 1;
    let columnIndex;
    if (childNumber % 3 === 1) columnIndex = 0;
    else if (childNumber % 3 === 2) columnIndex = 1;
    else columnIndex = 2;

    columnHeights[columnIndex] += totalHeight;
  });

  const maxColumnHeight = Math.max(...columnHeights);
  masonryContainer.style.height = maxColumnHeight + 'px';
}

window.addEventListener('load', setMasonryHeight);
window.addEventListener('resize', setMasonryHeight);
```

This function goes through the list of items within the masonry and calculates the height of each column. Then, we set the height of the container to the height of the tallest column.

**Note**: This implementation doesn't take the container's padding and border into consideration. ~~I'm too lazy to add it~~ This is left as an exercise for the reader.

## Backmatter

[^1]: The generalized version of this is that we need `n - 1`, where `n` is the number of columns we want in the masonry.
