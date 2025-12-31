---
title: Float Like Double Precision, Sting Like Square Root Computation
summary: My program for computing square roots broke for extremely large and small numbers. I ended up exploring how numbers are represented in computers as a part of my solution. I thought it would be fun to write this up and share with my study partner.
categories: programming, sicp
date: 2025-12-31T10:59:00+05:30
---

Recently, [Ajinkya](https://www.shindebhau.com/) and I revived our [SICP](https://github.com/shindeajinkya/SICP-26) [runs](https://github.com/nrabhiram/sicp-exercises). One of my [predictions for the coming year](/predictions-for-2026) is that I'll finish this book. Although ambitious, I think that this is doable for at least the first 3 chapters[^1] if I don't spend exorbitant amounts of time marveling at the beauty of the prose, as I did previously.

I promised myself that I wouldn't hyperfixate on things this time. And...and yet, I find myself breaking it once again. So early in the first chapter at that. How embarassing! But, in my defense, I came across a problem that required a bit of digging around that led to my unearthing some interesting insights that merited a dedicated write-up.

## The Problem With Computing Square Roots

The following program computes the square-root of the provided input value.

```racket
(define (square x) (* x x))

(define (good-enough? guess x)
  (< (abs (- (square guess) x))
     0.001))

(define (average x y)
  (/ (+ x y) 2))

(define (improve guess x)
  (average guess (/ x guess)))

(define (sqrt-iter guess x)
  (if (good-enough? guess x)
      guess
      (sqrt-iter (improve guess x) x)))

(define (sqrt x) (sqrt-iter 1.0 x))
```

But, this program breaks both when we're dealing with extremely small numbers and extremely large numbers. Why is that? And can we circumvent these problems?

## Small Numbers

The reasoning for small numbers is simple, so let's get that out of the way. For numbers that are less than the tolerance limit by at least an order of magnitude, we get inaccurate answers because the computation terminates early. We quickly end up with a guess that satisfies the `good-enough?` predicate. But, our tolerance value has been hard-coded. Because this value has been prescribed, it doesn't respect the scale of the radicand[^2] when it's extremely small (or extremely large, for that matter, but we'll get to that later). Although the guess could be further improved, the computation stops because the square of the guess falls within the tolerance limit of the radicand.

```racket
(sqrt 0.000001)
0.031260655525445276
(square (sqrt 0.000001))
0.0009772285838805523
```

But, this isn't right. This is analogous to giving the width of a coin ±0.1km. The precision is significantly larger than the values we're dealing with. This leads to inaccurate results.

## Precision and Accuracy

The difference between these two terms is subtle. It's easy to trip up and use them interchangeably, so, before proceeding, it's worth distinguishing the quantities that they represent.

- **accuracy**: how close the guess is to the actual value
- **precision**: the number of digits to which a guess can be specified

To substantiate this with an example, in the case of tiny numbers, we yield inaccurate results, i.e. results that aren't close to the actual value, because our tolerance limit lacks sufficient precision.

## Approximations

Now, this is where things get really interesting. The crux of our problem with extremely large numbers (and really small numbers) can be summarized as shown:

> The number of real numbers we can represent is infinite, but the number of bits we have to represent these values is finite.

For example, all of the following are valid real numbers:

- 0.001
- 0.0000001
- 0.000000000001
- 0.00000000000000000001
- 123456789123456789123456789
- 9999999999999999999999999999
- and so on

But, we can't accurately represent all of these numbers. We have a predetermined number of bits to represent the information about this number. When dealing with such extreme numbers, we might run out of bits to precisely represent them. So, the computer rounds it to the closest number it can represent.

## How Double Precision Works

The computer uses the `float` data type to represent real numbers. There are mainly two `float` types: 

- single precision, i.e. `float32`
- double precision, i.e. `float64`

These two representations differ in the number of bits they allocate for storing a number (`32` and `64` respectively). They work in similar ways, but you can represent a wider range of numbers with double precision.

In double precision, the format allocates:

- 1 bit for the sign, i.e. `+` or `-`,
- 52 bits for the mantissa, with an implicit leading bit
- 11 bits for the exponent

### What’s a Mantissa?

The mantissa, aka the significand, stands for the fractional part of a number represented in scientific notation. For example, we represent `1234` as `1.234 × 10⁴`. In this case, `1.234` is referred to as the mantissa. This is the base-10 notation — in `1234`, `4` is in the *1's* place, `3` is in the *10's* place, `2` is in the *100's* place, and `1` is in the *1000's* place. When we represent a number in scientific notation, the mantissa ranges between `1` (inclusive) and `10` (exclusive). If the mantissa's greater than or equal to `10`, it means that the scale, i.e. the exponent, can be increased by an order of magnitude.

Similarly, we can also represent a number in scientific notation using the base-2 format (binary). In base-2, the first digit is the *1's* (2⁰) place, the second digit is the *2's* (2¹) place, the third digit is the *4's* (2²) place, and so on. `1234` in binary is `10011010010`. When represented in scientific notation, we get `1.0011010010 × 2¹⁰`. Notice how in base-2, the fractional part can only exist between `1` (inclusive) and `2` (exclusive).

### Why Does It Have an Implicit Leading Bit?

For any number that's represented in base-2, the whole number part of the significand is always `1`. This means that the leading `1` will always exist (for non-zero values), and we need not assign a separate bit for storing this. This also means that we have 1 free bit of precision. Yippee! Wahoo!

### How Are Exponents Represented?

Since we have 11 bits to represent the exponent (`b × 2ⁿ`), this means that we can represent 2048 numbers. We want to accommodate negative exponent values as well. So, we pick the middle point as our effective 0. This means that we can represent exponent values ranging from `-1023` to `1024`.

So, in `float64`, the mantissa for `1234` is represented as `.0011010010` and the exponent is the exponent is `2¹⁰` (stored as `1033` after applying the bias of `1023`).

### Special Cases

There are a few special cases that we haven't covered in the double precision format of representing numbers. If the mantissa is always between `1` and `2`, how can we represent `0`? And what about really large values that tend to infinity? And what if we perform an invalid operation with numbers, such as dividing `0` by `0`? What then? Huh, huh?

The spec for double precision states that we reserve the 2 extreme values that can be represented by the exponent bits for these special cases, i.e. `0` and `2047`. This means that we can only represent exponent values from `-1022` to `1023`.

Now, here are the rules for representing these special values:

- `±0`: a mantissa of value `0` and an exponent of value `0`
- `±∞` (infinity): a mantissa of value `0` and an exponent of value `2047`
- `NaN` (not a number, i.e. invalid operations): a mantissa of a non-zero value and an exponent of value `2047`

## When Do Floats Break?

The amount of precision we have when it comes to extremely large (and extremely small numbers) can be calculated approximately as shown.

```
10ᵃ = 2⁵³
log₁₀ 10ᵃ = log₁₀ 2⁵³
a × log₁₀ 10  = 53 × log₁₀ 2
a = 53 × 0.3010
a = 15.953
```

So, in the 53 bits available for the mantissa (which determine the degree of precision with which we can measure a number, be it in base-2 or base-10), we can represent numbers with 15-17 digits of precision. If a number has a degree of precision that exceeds this, we have to round the number back to the closest value that the computer can represent. This can lead to inaccurate results. For numbers that are, say, 20 digits long, we end up with significantly larger approximation errors.

For example, `12345678912345678912`, when represented in `float64`, gives `12345678912345679872`. This value is off by `128`.

## Large Numbers

In the case of extremely large numbers, we're thrown into an infinite loop and the computation never ends. We reach a point where the guess can't be improved further, i.e. it converges. In some cases, if we're lucky, the absolute difference between the square of our best guess and the radicand is exactly `0`. But, in most cases, it exceeds the tolerance. When we try to improve the guess, the value that we end up with typically exists between two consecutive numbers that can be represented in double precision, and it gets rounded back to the previous guess.

## The Solution: Relative Precision

These approximation errors between the actual square-root of a number and our best guess may seem huge if viewed in isolation. But, relative to the magnitude of the number itself, it's not that big of a deal. We can fix this by using a relative tolerance instead of an absolute one.

```racket
(define (square x) (* x x))

(define (good-enough? guess prev-guess)
  (< (abs (/ (- guess prev-guess) prev-guess)) 0.00000000001))

(define (average x y)
  (/ (+ x y) 2))

(define (improve guess x)
  (average guess (/ x guess)))

(define (sqrt-iter guess x)
  (if (good-enough? (improve guess x) guess)
      guess
      (sqrt-iter (improve guess x) x)))

(define (sqrt x) (sqrt-iter 1.0 x))
```

Instead of using a fixed tolerance value, we can check whether the guess changes significantly relative to its current value.

```racket
(sqrt 0.000001)
0.0010000000000000117
(square (sqrt 0.000001))
1.0000000000000235e-6
(sqrt 123456789012345)
11111111.061111081
(square (sqrt 123456789012345))
123456789012345.02
```

## Concluding Thoughts

In my previous attempt, I was okay with accepting that large numbers had to be approximated by the computer. However, in this run, I wasn't satiated by this answer. I wanted to spend some effort to understand how numbers are represented in floating point, and how `float64` numbers can hold up to 16 digits of precision. I'm sure there's a lot more that goes on under the hood, but I've come out of this with a greater appreciation for how much thought has been put into creating a standard for what is now a primitive in most popular programming languages.

## Additional Resources

- [SICP solution for Exercise 1.7](https://sicp-solutions.net/post/sicp-solution-exercise-1-7/)
- [How Floating-Point Numbers Are Represented](https://youtu.be/bbkcEiUjehk?si=F-ptlgszCUsQmERU)

## Backmatter

[^1]: I've heard that it takes a lot of intentional effort to wrap your head around the ideas introduced in the last couple of chapters.

[^2]: The number for which we're computing the square root.
