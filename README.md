# Notepad 2 - Ken's Chronicle

### Challenge Description
"Years ago, when I was backpacking across Western Europe, I was just outside of Barcelona, hiking in the foothills of Mount Tibidabo. I was at the end of this path, and I came to a clearing, and there was a lake, very secluded, and there were tall trees all around. It was dead silent. Gorgeous. And across the lake, I saw a beautiful woman bathing herself. But she was crying...". There is a rumor that the rest of the story is on this website. Help Ken Adams uncover it.  

### Challenge File
[Source](./Handout/kenschronicles.zip)

### MD5 Hash:
`037a5f449a5925aaf0aa93b3bfb5a8584a36c9dd`

### Short Writeup

* Use `Cross-Origin-Opener-Policy:same-origin` header as a way to identify if your startsWith character was wrong
* If the header was set, and the page was opened using `x = window.open()`, `x.closed` would be true, else false
* Refer [exploit](./Admin/exploit)

### Author
[Az3z3l](https://twitter.com/Az3z3l)

### Flag
`inctf{tis_a_mooo_point_lkbsdgbyhsl}`