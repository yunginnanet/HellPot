package main

// Src is A. A. Milne's "Once on a Time" as transcribed by Project Gutenberg.
const Src = `<H1 ALIGN="center">
ONCE ON A TIME
</H1>

<H3 ALIGN="center">
<i>By</i>
</H3>

<H2 ALIGN="center">
A.A. Milne
</H2>

<BR><BR><BR>

<H3 ALIGN="center">
DECORATED<BR>
BY CHARLES<BR>
ROBINSON
</H3>

<BR>

<H4 ALIGN="center">
GROSSET & DUNLAP<BR>
Publishers&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;New York<BR>
By Arrangement with G. P. Putnam's Sons
</H4>

<br>
<p class="noindent" align="center"><img src="images/0008X.jpg" alt="[Illustration: Copyright page decoration of a dark-haired girl in medieval garb]"></p>

<br>

<H5 ALIGN="center">
Copyright, 1922
<BR>
by
<BR>
A. A. Milne
</H5>

<br>

<br>
<p class="noindent" align="center"><img src="images/0009X.jpg" alt="[Illustration: A young blonde girl in medieval garb]"></p>

<H3 ALIGN="center">
PREFACE
</H3>

<P>
This book was written in 1915, for the amusement of my wife and myself
at a time when life was not very amusing; it was published at the end
of 1917; was reviewed, if at all, as one of a parcel, by some brisk
uncle from the Tiny Tots Department; and died quietly, without
seriously detracting from the interest which was being taken in the
World War, then in progress.
</P>

<P>
It may be that the circumstances in which the book was written have
made me unduly fond of it.  When, as sometimes happens, I am
introduced to a stranger who starts the conversation on the right
lines by praising, however insincerely, my books, I always say, "But
you have not read the best one." Nine times out of ten it is so.  The
tenth takes a place in the family calendar; St. Michael or St. Agatha,
as the case may be, a red-letter or black-letter saint, according to
whether the book was bought or borrowed.  But there are few such
saints, and both my publisher and I have the feeling (so common to
publishers and authors) that there ought to be more.  So here comes
the book again, in a new dress, with new decorations, yet much, as far
as I am concerned, the same book, making the same appeal to me; but,
let us hope, a new appeal, this time, to others.
</P>

<P>
For whom, then, is the book intended?  That is the trouble.  Unless I
can say, "For those, young or old, who like the things which I like,"
I find it difficult to answer.  Is it a children's book?  Well, what
do we mean by that?  Is <i>The Wind in the Willows</i> a children's book?
Is <i>Alice in Wonderland?</i>  Is <i>Treasure Island?</i>  These are
masterpieces which we read with pleasure as children, but with how
much more pleasure when we are grown-up.  In any case what do we mean
by "children"?  A boy of three, a girl of six, a boy of ten, a girl of
fourteen&mdash;are they all to like the same thing?  And is a book
"suitable for a boy of twelve" any more likely to please a boy of
twelve than a modern novel is likely to please a man of thirty-seven;
even if the novel be described truly as "suitable for a man of
thirty-seven"?  I confess that I cannot grapple with these difficult
problems.
</P>

<P>
But I am very sure of this: that no one can write a book which
children will like, unless he write it for himself first.  That being
so, I shall say boldly that this is a story for grown-ups.  How
grown-up I did not realise until I received a letter from an unknown
reader a few weeks after its first publication; a letter which said
that he was delighted with my clever satires of the Kaiser, Mr. Lloyd
George and Mr. Asquith, but he could not be sure which of the
characters were meant to be Mr. Winston Churchill and Mr. Bonar Law.
Would I tell him on the enclosed postcard?  I replied that they were
thinly disguised on the title-page as Messrs. Hodder & Stoughton.  In
fact, it is not that sort of book.
</P>

<P>
But, as you see, I am still finding it difficult to explain just what
sort of book it is.  Perhaps no explanation is necessary.  Read in it
what you like; read it to whomever you like; be of what age you like;
it can only fall into one of two classes.  Either you will enjoy it,
or you won't.
</P>

<P>
It is that sort of book.
</P>

<P>
A. A. Milne.
</P>

<br><br>

<p class="noindent" align="center"><img src="images/0013X.jpg" alt="[Illustration: Contents page decoration of a child bent over a large boot]"></p>

<br><br>

<H3 ALIGN="center">
CONTENTS
</H3>

<P>
<a href="#chap01">I.&mdash;The King of Euralia has a Visitor to Breakfast</a>
</P>

<P>
<a href="#chap02">II.&mdash;The Chancellor of Barodia has a Long Walk Home</a>
</P>

<P>
<a href="#chap03">III.&mdash;The King of Euralia Draws his Sword</a>
</P>

<P>
<a href="#chap04">IV.&mdash;The Princess Hyacinth Leaves it to the Countess</a>
</P>

<P>
<a href="#chap05">V.&mdash;Belvane Indulges her Hobby</a>
</P>

<P>
<a href="#chap06">VI.&mdash;There are no Wizards in Barodia</a>
</P>

<P>
<a href="#chap07">VII.&mdash;The Princess Receives a Letter and Writes One</a>
</P>

<P>
<a href="#chap08">VIII.&mdash;Prince Udo Sleeps Badly</a>
</P>

<P>
<a href="#chap09">IX.&mdash;They are Afraid of Udo</a>
</P>

<P>
<a href="#chap10">X.&mdash;Charlotte Patacake Astonishes the Critics</a>
</P>

<P>
<a href="#chap11">XI.&mdash;Watercress Seems to go with the Ears</a>
</P>

<P>
<a href="#chap12">XII.&mdash;We Decide to Write to Udo's Father</a>
</P>

<P>
<a href="#chap13">XIII.&mdash;"Pink" Rhymes with "Think"</a>
</P>

<P>
<a href="#chap14">XIV.&mdash;"Why Can't you be like Wiggs?"</a>
</P>

<P>
<a href="#chap15">XV.&mdash;There is a Lover Waiting for Hyacinth</a>
</P>

<P>
<a href="#chap16">XVI.&mdash;Belvane Enjoys Herself</a>
</P>

<P>
<a href="#chap17">XVII.&mdash;The King of Barodia Drops the Whisker Habit</a>
</P>

<P>
<a href="#chap18">XVIII.&mdash;The Veteran of the Forest Entertains Two Very Young People</a>
</P>

<P>
<a href="#chap19">XIX.&mdash;Udo Behaves Like a Gentleman</a>
</P>

<P>
<a href="#chap20">XX.&mdash;Coronel Knows a Good Story when he Hears it</a>
</P>

<P>
<a href="#chap21">XXI.&mdash;A Serpent Coming after Udo</a>
</P>

<P>
<a href="#chap22">XXII.&mdash;The Seventeen Volumes go back Again</a>
</P>
<br><br>
<p class="noindent" align="center">
<img src="images/0015X.jpg" alt="[Illustration: A dark-haired girl in medieval garb in a pastoral scene]"></p>

<BR><BR>

<H3 ALIGN="center">
ILLUSTRATIONS
</H3>

<P>
<a href="#img0020">
A Map of Euralia showing the Adjacent Country of Barodia and the
far-distant Araby
</a>
</P>

<P>
<a href="#img0021X">
He was a Man of Simple Tastes
</a>
</P>

<P>
<a href="#img0026">
"Most extraordinary," said the King</a>
</P>

<P>
<a href="#img0046">
He found the King nursing a Bent Whisker and in the very Vilest of Tempers
</a>
</P>

<P>
<a href="#img0060">
"Try it on me," cried the Countess</a>
</P>

<P>
<a href="#img0078">
Five Times he had come back to give her his Last Instructions</a>
</P>

<P>
<a href="#img0102">
Armed to the Teeth, Amazon after Amazon marched by</a>
</P>

<P>
<a href="#img0118">
When the Respective Armies returned to Camp they found Their Majesties
asleep</a>
</P>

<P>
<a href="#img0132">
The Rabbit was gone, and there was a Fairy in front of her</a>
</P>

<P>
<a href="#img0154">
As Evening fell they came to a Woodman's Cottage at the Foot of a High
Hill</a>
</P>

<P>
<a href="#img0168">
"Coronel, here I am," said Udo pathetically, and he stepped out</a>
</P>

<P>
<a href="#img0186">
Twenty-one Minutes later Henrietta Crossbuns was acknowledging a Bag
of Gold
</a>
</P>

<P>
<a href="#img0200">
Princess Hyacinth gave a Shriek and faltered slowly backwards
</a>
</P>

<P>
<a href="#img0220">
"Now we can talk," said Hyacinth
</a>
</P>

<P>
<a href="#img0242">
He forgot his Manners, and made a Jump towards her
</a>
</P>

<P>
<a href="#img0243">
She glided gracefully behind the Sundial in a Pretty Affectation of
Alarm
</a>
</P>

<P>
<a href="#img0262">
When anybody of Superior Station or Age came into the Room she rose
and curtsied
</a>
</P>

<P>
<a href="#img0274">
And then she danced
</a>
</P>

<P>
<a href="#img0284">
"Good Morning," said Belvane
</a>
</P>

<P>
<a href="#img0308">
The Tent seemed to swim before his Eyes, and he knew no more
</a>
</P>

<P>
<a href="#img0332">
She turned round and went off daintily down the Hill
</a>
</P>

<P>
<a href="#img0352">
Let me present to you my friend the Duke Coronel
</a>
</P>

<P>
<a href="#img0368">
As the Towers of the Castle came in sight, Merriwig drew a Deep Breath
of Happiness
</a>
</P>

<P>
<a href="#img0396">
Belvane leading the Way with her Finger to her Lips
</a>
</P>

<P>
<a href="#img0397">
Merriwig following with an Exaggerated Caution
</a>
</P>

<P>
<a href="#img0412">
He was a Pleasant-looking Person, with a Round Clean-shaven Face
</a>
</P>

<P>
<a href="#img0420X">
Roger Scurvilegs
</a>
</P>
<p class="noindent" align="center"><img src="images/0017X.jpg" alt="Illustration: End of Illustration List Decoration"></p>

<br><br><br>

<p class="noindent" align="center">
<img src="images/0019.jpg" alt="Illustration: Page 1 Decoration">
</p>
<br><br>

<p class="noindent" align="center">
<a name="img0020"></a>
<img src="images/0020.jpg" alt="[Frontispiece: A Map of Euralia showing the Adjacent Country of Barodia and the far-distant Araby]"></p>

<BR><BR><BR>

<p class="noindent" align="center">
<a name="img0021X"></a>
<A NAME="chap01"></A>
<img src="images/0021X.jpg" alt="[Illustration: He was a Man of Simple Tastes]">
</p>


<H3 ALIGN="center">
CHAPTER I
</H3>

<H3 ALIGN="center">
THE KING OF EURALIA HAS A VISITOR TO BREAKFAST
</H3>


<P>
King Merriwig of Euralia sat at breakfast on his castle walls.  He
lifted the gold cover from the gold dish in front of him, selected a
trout and conveyed it carefully to his gold plate.  He was a man of
simple tastes, but when you have an aunt with the newly acquired gift
of turning anything she touches to gold, you must let her practise
sometimes.  In another age it might have been fretwork.
</P>

<P>
"Ah," said the King, "here you are, my dear."  He searched for his
napkin, but the Princess had already kissed him lightly on the top of
the head, and was sitting in her place opposite to him.
</P>

<P>
"Good morning, Father," she said; "I'm a little late, aren't I?  I've
been riding in the forest."
</P>

<P>
"Any adventures?" asked the King casually.
</P>

<P>
"Nothing, except it's a beautiful morning."
</P>

<P>
"Ah, well, perhaps the country isn't what it was.  Now when I was a
young man, you simply couldn't go into the forest without an adventure
of some sort.  The extraordinary things one encountered!  Witches,
giants, dwarfs&mdash;&mdash;.  It was there that I first met your mother," he
added thoughtfully.
</P>

<P>
"I wish I remembered my mother," said Hyacinth.
</P>

<P>
The King coughed and looked at her a little nervously.
</P>

<P>
"Seventeen years ago she died, Hyacinth, when you were only six months
old.  I have been wondering lately whether I haven't been a little
remiss in leaving you motherless so long."
</P>

<P>
The Princess looked puzzled.  "But it wasn't your fault, dear, that
mother died."
</P>

<P>
"Oh, no, no, I'm not saying that.  As you know, a dragon carried her
off and&mdash;well, there it was.  But supposing"&mdash;he looked at her
shyly&mdash;"I had married again."
</P>

<P>
The Princess was startled.
</P>

<P>
"Who?" she asked.
</P>

<P>
The King peered into his flagon.  "Well," he said, "there <i>are</i>
people."
</P>

<P>
"If it had been somebody <i>very</i> nice," said the Princess wistfully,
"it might have been rather lovely."
</P>

<P>
The King gazed earnestly at the outside of his flagon.
</P>

<P>
"Why 'might have been?'" he said.
</P>

<P>
The Princess was still puzzled.  "But I'm grown up," she said; "I
don't want a mother so much now."
</P>

<P>
The King turned his flagon round and studied the other side of it.
</P>

<P>
"A mother's&mdash;er&mdash;tender hand," he said, "is&mdash;er&mdash;never&mdash;&mdash;" and then
the outrageous thing happened.
</P>

<P>
It was all because of a birthday present to the King of Barodia, and
the present was nothing less than a pair of seven-league boots.  The
King being a busy man, it was a week or more before he had an
opportunity of trying those boots. Meanwhile he used to talk about
them at meals, and he would polish them up every night before he went
to bed.  When the great day came for the first trial of them to be
made, he took a patronising farewell of his wife and family, ignored
the many eager noses pressed against the upper windows of the Palace,
and sailed off.  The motion, as perhaps you know, is a little
disquieting at first, but one soon gets used to it.  After that it is
fascinating. He had gone some two thousand miles before he realised
that there might be a difficulty about finding his way back.  The
difficulty proved at least as great as he had anticipated. For the
rest of that day he toured backwards and forwards across the country;
and it was by the merest accident that a very angry King shot in
through an open pantry window in the early hours of the morning.  He
removed his boots and went softly to bed. . . .
</P>

<P>
It was, of course, a lesson to him.  He decided that in the future he
must proceed by a recognised route, sailing lightly from landmark to
landmark.  Such a route his Geographers prepared for him&mdash;an early
morning constitutional, of three hundred miles or so, to be taken ten
times before breakfast.  He gave himself a week in which to recover
his nerve and then started out on the first of them.
</P>

<P class="noindent" align="center">
<a name="img0026"></a>
<img src="images/0026.jpg"
alt="[Illustration: &quot;Most extraordinary,&quot; said the King, verso]">
<img src="images/0027.jpg"
alt="[Illustration: &quot;Most extraordinary,&quot; said the King, recto]">
</P>

<P>
Now the Kingdom of Euralia adjoined that of Barodia, but whereas
Barodia was a flat country, Euralia was a land of hills.  It was
natural then that the Court Geographers, in search of landmarks,
should have looked towards Euralia; and over Euralia accordingly,
about the time when cottage and castle alike were breakfasting, the
King of Barodia soared and dipped and soared and dipped again.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
"A mother's tender hand," said the King of Euralia,
"is&mdash;er&mdash;never&mdash;good gracious!  What's that?"
</P>

<P>
There was a sudden rush of air; something came for a moment between
his Majesty and the sun; and then all was quiet again.
</P>

<P>
"What was it?" asked Hyacinth, slightly alarmed.
</P>

<P>
"Most extraordinary," said the King.  "It left in my mind an
impression of ginger whiskers and large boots.  Do we know anybody
like that?"
</P>

<P>
"The King of Barodia," said Hyacinth, "has red whiskers, but I don't
know about his boots."
</P>

<P>
"But what could he have been doing up there?  Unless&mdash;&mdash;"
</P>

<P>
There was another rush of wind in the opposite direction; once more
the sun was obscured, and this time, plain for a moment for all to
see, appeared the rapidly dwindling back view of the King of Barodia
on his way home to breakfast.
</P>

<P>
Merriwig rose with dignity.
</P>

<P>
"You're quite right, Hyacinth," he said sternly; "it <i>was</i> the King of
Barodia."
</P>

<P>
Hyacinth looked troubled.
</P>

<P>
"He oughtn't to come over anybody's breakfast table quite so quickly
as that.  Ought he, Father?"
</P>

<P>
"A lamentable display of manners, my dear.  I shall withdraw now and
compose a stiff note to him.  The amenities must be observed."
</P>

<P>
Looking as severe as a naturally jovial face would permit him, and
wondering a little if he had pronounced "amenities" right, he strode
to the library.
</P>

<P>
The library was his Majesty's favourite apartment.  Here in the
mornings he would discuss affairs of state with his Chancellor, or
receive any distinguished visitors who were to come to his kingdom in
search of adventure.  Here in the afternoon, with a copy of <i>What to
say to a Wizard</i> or some such book taken at random from the shelves,
he would give himself up to meditation.
</P>

<P>
And it was the distinguished visitors of the morning who gave him most
to think about in the afternoon.  There were at this moment no fewer
than seven different Princes engaged upon seven different enterprises,
to whom, in the event of a successful conclusion, he had promised the
hand of Hyacinth and half his kingdom.  No wonder he felt that she
needed the guiding hand of a mother.
</P>

<P>
The stiff note to Barodia was not destined to be written.  He was
still hesitating between two different kinds of nib, when the door was
flung open and the fateful name of the Countess Belvane was announced.
</P>

<P>
The Countess Belvane!  What can I say which will bring home to you
that wonderful, terrible, fascinating woman? Mastered as she was by
overweening ambition, utterly unscrupulous in her methods of achieving
her purpose, none the less her adorable humanity betrayed itself in a
passion for diary-keeping and a devotion to the simpler forms of
lyrical verse.  That she is the villain of the piece I know well; in
his <i>Euralia Past and Present</i> the eminent historian, Roger
Scurvilegs, does not spare her; but that she had her great qualities I
should be the last to deny.
</P>

<P>
She had been writing poetry that morning, and she wore green.  She
always wore green when the Muse was upon her: a pleasing habit which,
whether as a warning or an inspiration, modern poets might do well to
imitate.  She carried an enormous diary under her arm; and in her mind
several alternative ways of putting down her reflections on her way to
the Palace.
</P>

<P>
"Good morning, dear Countess," said the King, rising only too gladly
from his nibs; "an early visit."
</P>

<P>
"You don't mind, your Majesty?" said the Countess anxiously.  "There
was a point in our conversation yesterday about which I was not quite
certain&mdash;&mdash;"
</P>

<P>
"What <i>were</i> we talking about yesterday?"
</P>

<P>
"Oh, your Majesty," said the Countess, "affairs of state," and she
gave him that wicked, innocent, impudent, and entirely scandalous look
which he never could resist, and you couldn't either for that matter.
</P>

<P>
"Affairs of state, of course," smiled the King.
</P>

<P>
"Why, I made a special note of it in my diary."
</P>

<P>
She laid down the enormous volume and turned lightly over the pages.
</P>

<P>
"Here we are! '<i>Thursday.</i>  His Majesty did me the honour to consult
me about the future of his daughter, the Princess Hyacinth.  Remained
to tea and was very&mdash;&mdash;'  I can't quite make this word out."
</P>

<P>
"Let <i>me</i> look," said the King, his rubicund face becoming yet more
rubicund.  "It looks like 'charming,'" he said casually.
</P>

<P>
"Fancy!" said Belvane.  "Fancy my writing that!  I put down just what
comes into my head at the time, you know."  She made a gesture with
her hand indicative of some one who puts down just what comes into her
head at the time, and returned to her diary.  "'Remained to tea, and
was very charming.  Mused afterwards on the mutability of life!'" She
looked up at him with wide-open eyes.  "I often muse when I'm alone,"
she said.
</P>

<P>
The King still hovered over the diary.
</P>

<P>
"Have you any more entries like&mdash;like that last one?  May I look?"
</P>

<P>
"Oh, your Majesty!  I'm afraid it's <i>quite</i> private."  She closed the
book quickly.
</P>

<P>
"I just thought I saw some poetry," said the King.
</P>

<P>
"Just a little ode to a favourite linnet.  It wouldn't interest your
Majesty."
</P>

<P>
"I adore poetry," said the King, who had himself written a rhymed
couplet which could be said either forwards or backwards, and in the
latter position was useful for removing enchantments.  According to
the eminent historian, Roger Scurvilegs, it had some vogue in Euralia
and went like this:
</P>

<P class="poem">
          "<i>Bo, boll, bill, bole.</i> <BR>
           &nbsp;<i>Wo, woll, will, wole.</i>"<BR>
</P>

<P>
A pleasing idea, temperately expressed.
</P>

<P>
The Countess, of course, was only pretending.  Really she was longing
to read it.  "It's quite a little thing," she said.
</P>

<P class=poem>
          "<i>Hail to thee, blithe linnet,</i><BR>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<i>Bird thou clearly art,</i><BR>
           &nbsp;<i>That from bush or in it</i><BR>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<i>Pourest thy full heart!</i><BR>
           &nbsp;<i>And leads the feathered choir in song</i><BR>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<i>Taking the treble part.</i>"<BR>
</P>

<P>
"Beautiful," said the King, and one must agree with him. Many years
after, another poet called Shelley plagiarised the idea, but handled
it in a more artificial, and, to my way of thinking, decidedly
inferior manner.
</P>

<P>
"Was it a real bird?" said the King.
</P>

<P>
"An old favourite."
</P>

<P>
"Was it pleased about it?"
</P>

<P>
"Alas, your Majesty, it died without hearing it."
</P>

<P>
"Poor bird!" said his Majesty; "I think it would have liked it."
</P>

<P>
Meanwhile Hyacinth, innocent of the nearness of a mother, remained on
the castle walls and tried to get on with her breakfast.  But she made
little progress with it.  After all, it <i>is</i> annoying continually to
look up from your bacon, or whatever it is, and see a foreign monarch
passing overhead.  Eighteen more times the King of Barodia took
Hyacinth in his stride.  At the end of the performance, feeling rather
giddy, she went down to her father.
</P>

<P>
She found him alone in the library, a foolish smile upon his face, but
no sign of a letter to Barodia in front of him.
</P>

<P>
"Have you sent the Note yet?" she asked.
</P>

<P>
"Note?  Note?" he said, bewildered, "what&mdash;oh, you mean the Stiff Note
to the King of Barodia?  I'm just planning it, my love.  The exact
shade of stiffness, combined with courtesy, is a little difficult to
hit."
</P>

<P>
"I shouldn't be too courteous," said Hyacinth; "he came over eighteen
more times after you'd gone."
</P>

<P>
"Eighteen, eighteen, eight&mdash;my dear, it's outrageous."
</P>

<P>
"I've never had such a crowded breakfast before."
</P>

<P>
"It's positively insulting, Hyacinth.  This is no occasion for Notes.
We will talk to him in a language that he will understand."
</P>

<P>
And he went out to speak to the Captain of his Archers.
</P>
<p class="noindent" align="center">
<img src="images/0037X.jpg" alt="[Illustration: Decoration of a pile of books]">
</p>


<BR><BR><BR>
<p class="noindent" align="center">
<A NAME="chap02"></A>
<img src="images/0039X.jpg"
alt="[Illustration: Selection from next two-page drawing]">
</p>

<H3 ALIGN="center">
CHAPTER II
</H3>

<H3 ALIGN="center">
THE CHANCELLOR OF BARODIA HAS A LONG WALK HOME
</H3>


<P>
Once more it was early morning on the castle walls.
</P>

<P>
The King sat at his breakfast table, a company of archers drawn up in
front of him.
</P>

<P>
"Now you all understand," he said.  "When the King of Baro&mdash;when a
certain&mdash;well, when I say 'when,' I want you all to fire your arrows
into the air.  You are to take no aim; you are just to shoot your
arrows upwards, and&mdash;er&mdash;I want to see who gets highest.  Should
anything&mdash;er&mdash;should anything brush up against them on their way&mdash;not
of course that it's likely&mdash;well, in that case&mdash;er&mdash;in that case
something will&mdash;er&mdash;brush up against them.  After all, what <i>should?</i>"
</P>

<P>
"Quite so, Sire," said the Captain, "or rather, not at all."
</P>

<P>
"Very well.  To your places."
</P>

<P>
Each archer fitted an arrow to his bow and took up his position.  A
look-out man had been posted.  Everything was ready.
</P>

<P>
The King was decidedly nervous.  He wandered from one archer to
another asking after this man's wife and family, praising the polish
on that man's quiver, or advising him to stand with his back a little
more to the sun.  Now and then he would hurry off to the look-out man
on a distant turret, point out Barodia on the horizon to him, and
hurry back again.
</P>

<P>
The look-out knew all about it.
</P>

<P>
"Royalty over," he bellowed suddenly.
</P>

<P>
"When!" roared the King, and a cloud of arrows shot into the air.
</P>

<P>
"Well done!" cried Hyacinth, clapping her hands.  "I mean, how could
you?  You might have hurt him."
</P>

<P>
"Hyacinth," said the King, turning suddenly; "you here?"
</P>

<P>
"I have just come up.  Did you hit him?"
</P>

<P>
"Hit who?"
</P>

<P>
"The King of Barodia, of course."
</P>

<P>
"The King of&mdash;&mdash;  My dear child, what could the King of Barodia be
doing here?  My archers were aiming at a hawk that they saw in the
distance."  He beckoned to the Captain.  "Did you hit that hawk?" he
asked.
</P>

<P>
"With one shot only, Sire.  In the whisk&mdash;in the tail feathers."
</P>

<P>
The King turned to Hyacinth.
</P>

<P>
"With one shot only in the whisk&mdash;in the tail feathers," he said.
"What was it, my dear, that you were saying about the King of
Barodia?"
</P>

<P>
"Oh, Father, you are bad.  You hit the poor man right in the whisker."
</P>

<P>
"His Majesty of Barodia!  And in the whisker!  My dear child, this is
terrible!  But what can he have been doing up there?  Dear, dear, this
is really most unfortunate.  I must compose a note of apology about
this."
</P>

<P>
"I should leave the first note to him," said Hyacinth.
</P>

<P>
"Yes, yes, you're right.  No doubt he will wish to explain how he came
to be there.  Just a moment, dear."
</P>

<P>
He went over to his archers, who were drawn up in line
again.
</P>

<P>
"You may take your men down now," he said to the Captain.
</P>

<P>
"Yes, your Majesty."
</P>

<P>
His Majesty looked quickly round the castle walls, and then leant
confidentially towards the Captain.
</P>

<P>
"Er&mdash;which was the man who&mdash;er"&mdash; he fingered his cheek&mdash;"er&mdash;quite
so.  The one on the left?  Ah, yes."  He went to the man on the left
and put a bag of gold into his hand.
</P>

<P>
"You have a very good style with the bow, my man.  Your wrist action
is excellent.  I have never seen an arrow go so high."
</P>

<P>
The company saluted and withdrew.  The King and Hyacinth sat down to
breakfast.
</P>

<P>
"A little mullet, my dear?" he said.
</P>

<P>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
The Hereditary Grand Chancellor of Barodia never forgot that morning,
nor did he allow his wife to forget it.  His opening, "That reminds
me, dear, of the day when&mdash;&mdash;" though the signal of departure for any
guests, allowed no escape for his family.  They had to have it.
</P>

<P>
And indeed it was a busy day for him.  Summoned to the Palace at nine
o'clock, he found the King nursing a bent whisker and in the very
vilest of tempers.  His Majesty was for war at once, the Chancellor
leant towards the Stiff Note.
</P>

<P>
"At least, your Majesty," he begged, "let me consult the precedents
first."
</P>

<P>
"There is no precedent," said the King coldly, "for such an outrage as
this."
</P>

<P>
"Not precisely, Sire; but similar unfortunate occurrences
have&mdash;occurred."
</P>

<P>
"It was worse than an occurrence."
</P>

<P>
"I should have said an outrage, your Majesty.  Your late lamented
grandfather was unfortunate enough to come beneath the spell of the
King of Araby, under which he was compelled&mdash;or perhaps I should say
preferred&mdash;to go about on his hands and knees for several weeks.  Your
Majesty may recall how the people in their great loyalty adopted a
similar mode of progression.  Now although your Majesty's case is not
precisely on all fours&mdash;&mdash;"
</P>

<P>
"Not at all on all fours," said the King coldly.
</P>

<P>
"An unfortunate metaphor; I should say that although your Majesty's
case is not parallel, the procedure adopted in your revered
grandfather's case&mdash;&mdash;"
</P>

<P>
"I don't care what <i>you</i> do with your whiskers; I don't care what
<i>anybody</i> does with his whiskers," said the King, still soothing his
own tenderly; "I want the King of Euralia's blood."  He looked round
the Court.  "To any one who will bring me the head of the King, I will
give the hand of my daughter in marriage."
</P>

<P>
There was a profound silence. . . .
</P>

<P>
"Which daughter?" said a cautious voice at last.
</P>

<P>
"The eldest," said the King.
</P>

<P>
There was another profound silence. . . .
</P>

<P class="noindent" align="center">
<a name="img0046"></a><img src="images/0046.jpg" alt="[Illustration: He found the King nursing a bent whisker and in the
very vilest of tempers, verso]"><img src="images/0047.jpg" alt="[Illustration: He found the King
nursing a bent whisker and in the very vilest of tempers, recto]">
</P>

<P>
"My suggestion, your Majesty," said the Chancellor, "is that for the
present there should be merely an exchange of Stiff Notes; and that
meanwhile we scour the kingdom for an enchanter who shall take some
pleasant revenge for us upon his Majesty of Euralia.  For instance,
Sire, a king whose head has been permanently fixed on upside-down
lacks somewhat of that regal dignity which alone can command the
respect of his subjects.  A couple of noses, again, placed at
different angles, so they cannot both be blown together&mdash;&mdash;"
</P>

<P>
"Yes, yes," said the King impatiently, "<i>I'll</i> think of the things, if
once you can find the enchanter.  But they are not so common nowadays.
Besides, enchanters are delicate things to work with.  They have a
habit of forgetting which side they are on."
</P>

<P>
The Chancellor's mouth drooped piteously.
</P>

<P>
"Well," said the King condescendingly, "I'll tell you what we'll do.
You may send <i>one</i> Stiff Note and then we will declare war."
</P>

<P>
"Thank you, your Majesty," said the Chancellor.
</P>

<P>
So the Stiff Note was dispatched.  It pointed out that his Majesty of
Barodia, while in the act of taking his early morning constitutional,
had been severely insulted by an arrow.  This arrow, though
fortunately avoiding the more vital parts of his Majesty's person,
went so far as to wound a favourite whisker.  For this the fullest
reparation must be made . . . and so forth and so on.
</P>

<P>
Euralia's reply was not long delayed.  It expressed the deepest
concern at the unhappy accident which had overtaken a friendly
monarch.  On the morning in question, his Majesty had been testing his
archers in a shooting competition at a distant hawk; which
competition, it might interest his Majesty of Barodia to know, had
been won by Henry Smallnose, a bowman of considerable promise.  In the
course of the competition it was noticed that a foreign body of some
sort brushed up against one of the arrows, but as this in no way
affected the final placing of the competitors, little attention was
paid to it.  His Majesty of Barodia might rest assured that the King
had no wish to pursue the matter farther.  Indeed, he was always glad
to welcome his Barodian Majesty on these occasions.  Other shooting
competitions would be arranged from time to time, and if his Majesty
happened to be passing at the moment, the King of Euralia hoped that
he would come down and join them. Trusting that her Majesty and their
Royal Highnesses were well, . . . and so on and so forth.
</P>

<P>
The Grand Chancellor of Barodia read this answer to his Stiff Note
with a growing feeling of uneasiness.  It was he who had exposed his
Majesty to this fresh insult; and, unless he could soften it in some
way, his morning at the Palace might be a painful one.
</P>

<P>
As he entered the precincts, he wondered whether the King would be
wearing the famous boots, and whether they kicked seven leagues as
easily as they strode them.  He felt more and more that there were
notes which you could break gently, and notes which you
couldn't. . . .
</P>

<P>
Five minutes later, as he started on his twenty-one mile walk home, he
realised that this was one of the ones which you couldn't.
</P>

<P>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
This, then, was the real reason of the war between Euralia and
Barodia.  I am aware that in saying this I differ from the eminent
historian, Roger Scurvilegs.  In Chapter IX of his immortal work,
<i>Euralia Past and Present</i>, he attributes the quarrel between the two
countries to quite other causes.  The King of Barodia, he says,
demanded the hand of the Princess Hyacinth for his eldest son.  The
King of Euralia made some commonplace condition as that his Royal
Highness should first ride his horse up a glassy mountain in the
district, a condition which his Majesty of Barodia strongly resented.
I am afraid that Roger is incurably romantic; I have had to speak to
him about it before. There was nothing of the sentimental in the whole
business, and the facts are exactly as I have narrated them.
</P>
<p class="noindent" align="center"><img src="images/0052X.jpg" alt="[Illustration: End of chapter decoration.  Looks to be Belvane reading her diary, but it is very small.]">




<BR><BR><BR>
<p class="noindent" align="center">
<A NAME="chap03"></A>
<img src="images/0053X.jpg" alt="[Illustration: Detail of Merriwig from next large drawing]">
</p>

<H3 ALIGN="center">
CHAPTER III
</H3>

<H3 ALIGN="center">
THE KING OF EURALIA DRAWS HIS SWORD
</H3>

<P>
No doubt you have already guessed that it was the Countess Belvane who
dictated the King of Euralia's answer.  Left to himself, Merriwig
would have said, "Serve you jolly well right for stalking over my
kingdom."  His repartee was never very subtle.  Hyacinth would have
said, "Of course we're <i>awfully</i> sorry, but a whisker isn't <i>very</i>
bad, is it? and you really <i>oughtn't</i> to come to breakfast without
being asked."  The Chancellor would have scratched his head for a long
time, and then said, "Referring to Chap VII, Para 259 of the <i>King's
Regulations</i> we notice . . ."
</P>

<P>
But Belvane had her own way of doing things; and if you suggest that
she wanted to make Barodia's declaration of war inevitable, well, the
story will show whether you are right in supposing that she had her
reasons.  It came a little hard on the Chancellor of Barodia, but the
innocent must needs suffer for the ambitions of the unprincipled&mdash;a
maxim I borrow from <i>Euralia Past and Present;</i> Roger in his moral
vein.
</P>

<P>
"Well," said Merriwig to the Countess, "that's done it."
</P>

<P>
"It really is war?" asked Belvane.
</P>

<P>
"It is.  Hyacinth is looking out my armour at this moment."
</P>

<P>
"What did the King of Barodia say?"
</P>

<P>
"He didn't <i>say</i> anything.  He wrote 'W A R' in red on a dirty bit of
paper, pinned it to my messenger's ear, and sent him back again."
</P>

<P>
"How very crude," said the Countess.
</P>

<P>
"Oh, I thought it was&mdash;er&mdash;rather forcible," said the King awkwardly.
Secretly he had admired it a good deal and wished that he had been the
one to do it.
</P>

<P>
"Of course," said the Countess, with a charming smile, "that sort of
thing depends so <i>very</i> much on who does it.  Now from your Majesty it
would have seemed&mdash;dignified."
</P>

<P>
"He must have been very angry," said the King, picking up first one
and then another of a number of swords which lay in front of him.  "I
wish I had seen his face when he got my Note."
</P>

<P>
"So do I," sighed the Countess.  She wished it much more than the
King.  It is the tragedy of writing a good letter that you cannot be
there when it is opened: a maxim of my own, the thought never having
occurred to Roger Scurvilegs, who was a dull correspondent.
</P>

<P>
The King was still taking up and putting down his swords.
</P>

<P>
"It's very awkward," he muttered; "I wonder if Hyacinth&mdash;&mdash;"  He went
to the door and called "Hyacinth!"
</P>

<P>
"Coming, Father," called back Hyacinth, from a higher floor.
</P>

<P>
The Countess rose and curtsied deeply.
</P>

<P>
"Good morning, your Royal Highness."
</P>

<P>
"Good morning, Countess," said Hyacinth brightly.  She liked the
Countess (you couldn't help it), but rather wished she didn't.
</P>

<P>
"Oh, Hyacinth," said the King, "come and tell me about these swords.
Which is my magic one?"
</P>

<P>
Hyacinth looked at him blankly.
</P>

<P>
"Oh, Father," she said.  "I don't know at all.  Does it matter very
much?"
</P>

<P>
"My dear child, of course it matters.  Supposing I am fighting the
King of Barodia and I have my magic sword, then I'm bound to win.
Supposing I haven't, then I'm not bound to."
</P>

<P>
"Supposing you both had magic swords," said Belvane.  It was the sort
of thing she <i>would</i> say.
</P>

<P>
The King looked up slowly at her and began to revolve the idea in his
mind.
</P>

<P>
"Well, really," he said, "I hadn't thought of that.  Upon my word,
I&mdash;&mdash;"  He turned to his daughter.  "Hyacinth, what would happen if we
both had magic swords?"
</P>

<P>
"I suppose you'd go on fighting for ever," said Hyacinth.
</P>

<P>
"Or until the magic wore out of one of them," said Belvane innocently.
</P>

<P>
"There must be something about it somewhere," said the King, whose
morning was in danger of being quite spoilt by this new suggestion;
"I'd ask the Chancellor to look it up, only he's so busy just now."
</P>

<P>
"He'd have plenty of time while the combat was going on," said Belvane
thoughtfully.  Wonderful creature! she saw already the Chancellor
hurrying up to announce that the King of Euralia had won, at the very
moment when he lay stretched on the ground by a mortal thrust from his
adversary.
</P>

<P>
The King turned to his swords again.
</P>

<P>
"Well, anyway, I'm going to be sure of <i>mine</i>," he said. "Hyacinth,
haven't you <i>any</i> idea which it is?"  He added in rather a hurt voice,
"Naturally I left the marking of my swords to <i>you</i>."
</P>

<P>
His daughter examined the swords one by one.
</P>

<P>
"Here it is," she cried.  "It's got 'M' on it for 'magic.'"
</P>

<P>
"Or 'Merriwig,'" said the Countess to her diary.
</P>

<P>
The expression of joy on the King's face at his daughter's discovery
had just time to appear and fade away again.
</P>

<P>
"You are not being very helpful this morning, Countess," he said
severely.
</P>

<P>
Instantly the Countess was on her feet, her diary thrown to the
floor&mdash;no, never thrown&mdash;laid gently on the floor, and herself, hands
clasped at her breast, a figure of reproachful penitence before him.
</P>

<P>
"Oh, your Majesty, forgive me&mdash;if your Majesty had only asked me&mdash;I
didn't know your Majesty wanted me&mdash;I thought her Royal Highness&mdash;&mdash;
But <i>of course</i> I'll find your Majesty's sword for you."  Did she
stroke his head as she said this?  I have often wondered.  It would be
like her impudence, and her motherliness, and her&mdash;-and, in fact, like
her.  <i>Euralia Past and Present</i> is silent upon the point.  Roger
Scurvilegs, who had only seen Belvane at the unimpressionable age of
two, would have had it against her if he could, so perhaps there is
nothing in it.
</P>

<P>
"There!" she said, and she picked out the magic sword almost at once.
</P>

<P class="noindent" align="center">
<a name="img0060"></a><img src="images/0060.jpg" alt="[Illustration: &quot;Try it on me,&quot; cried the Countess, verso]"><img src="images/0061.jpg" alt="[Illustration: &quot;Try it on me,&quot; cried the Countess, recto]">
</P>

<P>
"Then I'll get back to my work," said Hyacinth cheerfully, and left
them to each other.
</P>

<P>
The King, smiling happily, girded on his sword.  But a sudden doubt
assailed him.
</P>

<P>
"Are you sure it's the one?"
</P>

<P>
"Try it on <i>me</i>," cried the Countess superbly, falling on her knees
and stretching up her arms to him.  The toe of her little shoe touched
her diary; its presence there uplifted her.  Even as she knelt she saw
herself describing the scene.  How do you spell "offered"? she
wondered.
</P>

<P>
I think the King was already in love with her, though he found it so
difficult to say the decisive words.  But even so he could only have
been in love a week or two; a fortnight in the last forty years; and
he had worn a sword since he was twelve.  In a crisis it is the old
love and not the greater love which wins (Roger's, but I think I agree
with him), and instinctively the King drew his sword.  If it were
magic a scratch would kill.  Now he would know.
</P>

<P>
Her enemies said that the Countess could not go pale; she had her
faults, but this was not one of them.  She whitened as she saw the
King standing over her with drawn sword.  A hundred thoughts chased
each other through her mind.  She wondered if the King would be sorry
afterwards; she wondered what the minstrels would sing of her, and if
her diary would ever be made public; most of all she wondered why she
had been such a fool, such a melodramatic fool.
</P>

<P>
The King came to himself with a sudden start.  Looking slightly
ashamed he put his sword back in its scabbard, coughed once or twice
to cover his confusion, and held his hand out to the Countess to
assist her to rise.
</P>

<P>
"Don't be absurd, Countess," he said.  "As if we could spare you at a
time like this.  Sit down and let us talk matters over seriously."
</P>

<P>
A trifle bewildered by the emotions she had gone through, Belvane sat
down, the beloved diary clasped tightly in her arms.  Life seemed
singularly sweet just then, the only drawback being that the minstrels
would not be singing about her after all.  Still, one cannot have
everything.
</P>

<P>
The King walked up and down the room as he talked.
</P>

<P>
"I am going away to fight," he said, "and I leave my dear daughter
behind.  In my absence, her Royal Highness will of course rule the
country.  I want her to feel that she can lean upon you, Countess, for
advice and support.  I know that I can trust you, for you have just
given me a great proof of your devotion and courage."
</P>

<P>
"Oh, your Majesty!" said Belvane deprecatingly, but feeling very glad
that it hadn't been wasted.
</P>

<P>
"Hyacinth is young and inexperienced.  She needs a&mdash;a&mdash;&mdash;"
</P>

<P>
"A mother's guiding hand," said Belvane softly.
</P>

<P>
The King started and looked away.  It was really too late to propose
now; he had so much to do before the morrow. Better leave it till he
came back from the war.
</P>

<P>
"You will have no official position," he went on hastily, "other than
your present one of Mistress of the Robes; but your influence on her
will be very great."
</P>

<P>
The Countess had already decided on this.  However there <i>is</i> a look
of modest resignation to an unsought duty which is suited to an
occasion of this kind, and the Countess had no difficulty in supplying
it.
</P>

<P>
"I will do all that I can, your Majesty, to help&mdash;gladly; but will not
the Chancellor&mdash;&mdash;"
</P>

<P>
"The Chancellor will come with me.  He is no fighter, but he is good
at spells."  He looked round to make sure that they were alone, and
then went on confidentially, "He tells me that he has discovered in
the archives of the palace a Backward Spell of great value.  Should he
be able to cast this upon the enemy at the first onslaught, he thinks
that our heroic army would have no difficulty in advancing."
</P>

<P>
"But there will be other learned men," said Belvane innocently, "so
much more accustomed to affairs than us poor women, so much better
able"&mdash;("What nonsense I'm talking," she said to herself)&mdash;"to advise
her Royal Highness&mdash;&mdash;"
</P>

<P>
"Men like that," said the King, "I shall want with me also.  If I am
to invade Barodia properly I shall need every man in the kingdom.
Euralia must be for the time a country of women only."  He turned to
her with a smile and said gallantly, "That will be&mdash;er&mdash;&mdash;  It
is&mdash;er&mdash;not&mdash;er&mdash;&mdash;.  One may well&mdash;er&mdash;&mdash;"
</P>

<P>
It was so obvious from his manner that something complimentary was
struggling to the surface of his mind, that Belvane felt it would be
kinder not to wait for it.
</P>

<P>
"Oh, your Majesty," she said, "you flatter my poor sex."
</P>

<P>
"Not at all," said the King, trying to remember what he had said.  He
held out his hand.  "Well, Countess, I have much to do."
</P>

<P>
"I, too, your Majesty."
</P>

<P>
She made him a deep curtsey and, clasping tightly the precious diary,
withdrew.
</P>

<P>
The King, who still seemed worried about something, returned to his
table and took up his pen.  Here Hyacinth discovered him ten minutes
later.  His table was covered with scraps of paper and, her eyes
lighting casually upon one of them, she read these remarkable words:
</P>

<P>
"<i>In such a land I should be a most contented subject.</i>"
</P>

<P>
She looked at some of the others.  They were even shorter:
</P>

<P>
"<i>That, dear Countess, would be my&mdash;&mdash;</i>"
</P>

<P>
"<i>A country in which even a King&mdash;&mdash;</i>"
</P>

<P>
"<i>Lucky country!</i>"
</P>

<P>
The last was crossed out and "<i>Bad</i>" written against it.
</P>

<P>
"Whatever are these, Father?" said Hyacinth.
</P>

<P>
The King jumped up in great confusion.
</P>

<P>
"Nothing, dear, nothing," he said.  "I was just&mdash;er&mdash;&mdash;  Of course I
shall have to address my people, and I was just jotting down a few&mdash;&mdash;
However, I shan't want them now." He swept them together, screwed
them up tight, and dropped them into a basket.
</P>

<P>
And what became of them? you ask.  Did they light the fires of the
Palace next morning?  Well, now, here's a curious thing.  In Chapter X
of <i>Euralia Past and Present</i> I happened across these words:
</P>

<P>
"The King and all the men of the land having left to fight the wicked
Barodians, Euralia was now a country of women only&mdash;<i>a country in
which even a King might be glad to be a subject</i>."
</P>

<P>
Now what does this mean?  Is it another example of literary theft?  I
have already had to expose Shelley.  Must I now drag into the light of
day a still worse plagiarism by Roger Scurvilegs?  The waste-paper
baskets of the Palace were no doubt open to him as to so many
historians.  But should he not have made acknowledgments?
</P>

<P>
I do not wish to be hard on Roger.  That I differ from him on many
points of historical fact has already been made plain, and will be
made still more plain as my story goes on.  But I have a respect for
the man; and on some matters, particularly those concerning Prince Udo
of Araby's first appearance in Euralia, I have to rely entirely upon
him for my information.  Moreover I have never hesitated to give him
credit for such of his epigrams as I have introduced into this book,
and I like to think that he would be equally punctilious to others.
We know his romantic way; no doubt the thought occurred to him
independently.  Let us put it at that, anyhow.
</P>

<P>
Belvane, meanwhile, was getting on.  The King had drawn his sword on
her and she had not flinched.  As a reward she was to be the power
behind the throne.
</P>

<P>
"Not necessarily <i>behind</i> the throne," said Belvane to herself.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap04"></A>
<img src="images/0071X.jpg" alt="[Illustration: Detail of Hyacinth on the castle walls]">
</p>

<H3 ALIGN="center">
CHAPTER IV
</H3>

<H3 ALIGN="center">
THE PRINCESS HYACINTH LEAVES IT TO THE COUNTESS
</H3>

<P>
It is now time to introduce Wiggs to you, and I find myself in a
difficulty at once.  What <i>was</i> Wiggs's position in the Palace?
</P>

<P>
This story is hard to tell, for I have to piece it together from the
narratives of others, and to supply any gaps in their stories from my
knowledge of how the different characters might be expected to act.
Perhaps, therefore, it is a good moment in which to introduce to you
the authorities upon whom I rely.
</P>

<P>
First and foremost, of course, comes Roger Scurvilegs.  His monumental
work, <i>Euralia Past and Present</i>, in seventeen volumes, towers upon my
desk as I write.  By the merest chance I picked it up (in a
metaphorical sense) at that little shop near&mdash;I forget its name, but
it's the third bookshop on the left as you come into London from the
New Barnet end.  Upon him I depend for the broad lines of my story,
and I have already indicated my opinion of the value of his work.
</P>

<P>
Secondly, come the many legends and ballads handed on to me years ago
by my aunt by marriage, one of the Cornish Smallnoses.  She claims to
be a direct descendant of that Henry Smallnose whose lucky shot
brought about the events which I am to describe.  I say she claims to
be, and one cannot doubt a lady's word in these matters; certainly she
used to speak about Henry with that mixture of pride and extreme
familiarity which comes best from a relation.  In all matters not
touching Henry, I feel that I can rely upon her; in its main lines her
narrative is strictly confirmed by Scurvilegs, and she brought to it a
picturesqueness and an appreciation of the true character of Belvane
which is lacking in the other; but her attitude towards Henry
Smallnose was absurd.  Indeed she would have had him the hero of the
story.  This makes Roger and myself smile.  We give him credit for the
first shot, and then we drop him.
</P>

<P>
Thirdly, Belvane herself.  Women like Belvane never die, and I met her
(or a reincarnation of her) at a country house in Shropshire last
summer.  I forget what she calls herself now, but I recognised her at
once; and, as I watched her, the centuries rolled away and she and I
were in Euralia, that pleasant country, together.  "Stayed to tea and
was very charming."  Would she have said that of me, I wonder?  But
I'm getting sentimental&mdash;Roger's great fault.
</P>

<P>
These then are my authorities; I consult them, and I ask myself, What
was Wiggs?
</P>

<P>
Roger speaks of her simply as an attendant upon the Princess.  Now we
know that the Princess was seventeen; Wiggs then would be about the
same age&mdash;a lady-in-waiting&mdash;perhaps even a little older.  Why not?
you say.  The Lady Wiggs, maid-of-honour to her Royal Highness the
Princess Hyacinth, eighteen and a bit, tall and stately.  Since she is
to endanger Belvane's plans, let her be something of a match for the
wicked woman.
</P>

<P>
Yes, but you would never talk like that if you had heard one of my
aunt's stories.  Nor if you had seen Belvane would you think that any
grown-up woman could be a match for her.
</P>

<P>
Wiggs was a child; I feel it in my bones.  In all the legends and
ballads handed down to me by my aunt she appears to me as a little
girl&mdash;Alice in a fairy story. Roger or no Roger I must have her a
child.
</P>

<P>
And even Roger cannot keep up the farce that she is a real
lady-in-waiting.  In one place he tells us that she dusts the throne
of the Princess; can you see her ladyship, eighteen last February,
doing that?  At other times he allows her to take orders from the
Countess; I ask you to imagine a maid-of-honour taking orders from any
but her own mistress.  Conceive her dignity!
</P>

<P>
A little friend, then, of Hyacinth's, let us say; ready to do anything
for anybody who loved, or appeared to love, her mistress.
</P>

<P>
The King had departed for the wars.  His magic sword girded to his
side, his cloak of darkness, not worn but rolled up behind him, lest
the absence of his usual extensive shadow should disturb his horse, he
rode at the head of his men to meet the enemy.  Hyacinth had seen him
off from the Palace steps.  Five times he had come back to give her
his last instructions, and a sixth time for his sword, but now he was
gone, and she was alone on the castle walls with Wiggs.
</P>

<P>
"Saying good-bye to fathers is very tiring," said Hyacinth.  "I do
hope he'll be all right.  Wiggs, although we oughtn't to mention it to
anybody, and although he's only just gone, we do think it will be
rather fun being Queen, don't we?"
</P>

<P>
"It must be lovely," said Wiggs, gazing at her with large eyes.  "Can
you really do whatever you like now?"
</P>

<P>
Hyacinth nodded.
</P>

<P>
"I always <i>did</i> whatever I liked," she said, "But now I really <i>can</i>
do it."
</P>

<P>
"Could you cut anybody's head off?"
</P>

<P>
"Easily," said the Princess confidently.
</P>

<P>
"I should hate to cut anybody's head off."
</P>

<P>
"So should I, Wiggs.  Let's decide to have no heads off just at
present&mdash;till we're more used to it."
</P>

<P>
Wiggs still kept her eyes fixed upon the Princess.
</P>

<P>
"Which is stronger," she asked, "you or a Fairy?"
</P>

<P>
"I knew you were going to ask something horrid like that," said
Hyacinth, pretending to be angry.  She looked quickly round to see
that nobody was listening, and then whispered in Wiggs's ear, "I am."
</P>

<P>
"O&mdash;oh!" said Wiggs.  "How lovely!"
</P>

<P>
"Isn't it?  Did you ever hear the story of Father and the Fairy?"
</P>

<P>
"His Majesty?"
</P>

<P>
"His Majesty the King of Euralia.  It happened in the forest one day
just after he became King."
</P>

<P>
Did <i>you</i> ever hear the story?  I expect not.  Well, then, you must
hear it.  But there will be too many inverted commas in it if I let
Hyacinth tell you, so I shall tell you myself.
</P>

<P class="noindent" align="center">
<a name="img0078"></a><img src="images/0078.jpg" alt="[Illustration: Five times he had come back to give her his last instructions, verso]">
<img src="images/0079.jpg" alt="[Illustration: Five times he had come back to give her his last instructions, recto]">
</P>

<P>
It was just after he became King.  He was so proud that he used to go
about saying, "I am the King.  I am the King." And sometimes, "The
King am I.  The King I am."  He was saying this one day in the forest
when a Fairy overheard him.  So she appeared in front of him and said,
"I believe you are the King?"
</P>

<P>
"I am the King," said Merriwig.  "I am the King, I am the&mdash;&mdash;"
</P>

<P>
"And yet," said the Fairy, "what is a King after all?"
</P>

<P>
"It is a very powerful thing to be a King," said Merriwig proudly.
</P>

<P>
"Supposing I were to turn you into a&mdash;a small sheep.  Then where would
you be?"
</P>

<P>
The King thought anxiously for a moment.
</P>

<P>
"I should like to be a small sheep," he said.
</P>

<P>
The Fairy waved her wand.
</P>

<P>
"Then you can be one," she said, "until you own that a Fairy is much
more powerful than a King."
</P>

<P>
So all at once he was a small sheep.
</P>

<P>
"Well?" said the Fairy.
</P>

<P>
"Well?" said the King.
</P>

<P>
"Which is more powerful, a King or a Fairy?"
</P>

<P>
"A King," said Merriwig.  "Besides being more woolly," he added.
</P>

<P>
There was silence for a little.  Merriwig began to eat some grass.
</P>

<P>
"I don't think much of Fairies," he said with his mouth full.  "I
don't think they're very powerful."
</P>

<P>
The Fairy looked at him angrily.
</P>

<P>
"They can't make you say things you don't want to say," he explained.
</P>

<P>
The Fairy stamped her foot.
</P>

<P>
"Be a toad," she said, waving her wand.  "A nasty, horrid, crawling
toad."
</P>

<P>
"I've <i>always</i> wanted&mdash;" began Merriwig&mdash;"to be a toad," he ended from
lower down.
</P>

<P>
"Well?" said the Fairy.
</P>

<P>
"I don't think much of Fairies," said the King.  "I don't think
they're very powerful."  He waited for the Fairy to look at him, but
she pretended to be thinking of something else.  After waiting a
minute or two, he added, "They can't make you say things you don't
want to say."
</P>

<P>
The Fairy stamped her foot still more angrily, and moved her wand a
third time.
</P>

<P>
"Be silent!" she commanded.  "And stay silent for ever!"
</P>

<P>
There was no sound in the forest.  The Fairy looked at the blue sky
through the green roof above her; she looked through the tall trunks
of the trees to the King's castle beyond; her eyes fell upon the
little glade on her left, upon the mossy bank on her right . . . but
she would not look down to the toad at her feet.
</P>

<P>
No, she wouldn't. . . .
</P>

<P>
She <i>wouldn't</i>. . . .
</P>

<P>
And yet&mdash;&mdash;
</P>

<P>
It was too much for her.  She could resist no longer.  She looked at
the nasty, horrid, crawling toad, the dumb toad at her feet that was
once a King.
</P>

<P>
And, catching her eye, the toad&mdash;<i>winked</i>.
</P>

<P>
Some winks are more expressive than others.  The Fairy knew quite well
what this one meant.  It meant:
</P>

<P>
"I don't think much of Fairies.  I don't think they're very powerful.
They can't make you say things you don't want to say."
</P>

<P>
The Fairy waved her wand in disgust.
</P>

<P>
"Oh, be a King again," she said impatiently, and vanished.
</P>

<P>
And so that is the story of how the King of Euralia met the Fairy in
the forest.  Roger Scurvilegs tells it well&mdash;indeed, almost as well as
I do&mdash;but he burdens it with a moral.  You must think it out for
yourself; I shall not give it to you.
</P>

<P>
Wiggs didn't bother about the moral.  Her elbows on her knees, her
chin resting on her hands, she gazed at the forest and imagined the
scene to herself.
</P>

<P>
"How wonderful to be a King like that!" she thought.
</P>

<P>
"That was a long time ago," explained Hyacinth.  "Father must have
been rather lovely in those days," she added.
</P>

<P>
"It was a very bad Fairy," said Wiggs.
</P>

<P>
"It was a very stupid one.  I wouldn't have given in to Father like
that."
</P>

<P>
"But there are good Fairies, aren't there?  I met one once."
</P>

<P>
"You, child?  Where?"
</P>

<P>
I don't know if it would have made any difference to Euralian history
if Wiggs had been allowed to tell about her Fairy then; as it was, she
didn't tell the story till later on, when Belvane happened to be near.
I regret to say that Belvane listened.  It was the sort of story that
<i>always</i> got overheard, she explained afterwards, as if that were any
excuse.  On this occasion she was just too early to overhear, but in
time to prevent the story being told without her.
</P>

<P>
"The Countess Belvane," said an attendant, and her ladyship made a
superb entry.
</P>

<P>
"Good morning, Countess," said Hyacinth.
</P>

<P>
"Good morning, your Royal Highness.  Ah, Wiggs, sweet child," she
added carelessly, putting out a hand to pat the sweet child's head,
but missing it.
</P>

<P>
"Wiggs was just telling me a story," said the Princess.
</P>

<P>
"Sweet child," said Belvane, feeling vaguely for her with the other
hand.  "<i>Could</i> I interrupt the story with a little business, your
Royal Highness?"
</P>

<P>
At a nod from the Princess, Wiggs withdrew.
</P>

<P>
"Well?" said Hyacinth nervously.
</P>

<P>
Belvane had always a curious effect on the Princess when they were
alone together.  There was something about her large manner which made
Hyacinth feel like a schoolgirl who has been behaving badly: alarmed
and apologetic.  I feel like this myself when I have an interview with
my publishers, and Roger Scurvilegs (upon the same subject) drags in a
certain uncle of his before whom (so he says) he always appears at his
worst.  It is a common experience.
</P>

<P>
"Just one or two little schemes to submit to your Majesty," said the
Countess.  "How silly of me&mdash;I mean, your Royal Highness.  Of course
your Royal Highness may not like them at all, but in case your Royal
Highness did, I just&mdash;well, I just wrote them out."
</P>

<P>
She unfolded, one by one, a series of ornamental parchments.
</P>

<P>
"They are beautifully written," said the Princess.
</P>

<P>
Belvane blushed at the compliment.  She had a passion for coloured
inks and rulers.  In her diary the day of the week was always
underlined in red, the important words in the day's doings being
frequently picked out in gold.  On taking up the diary you saw at once
that you were in the presence of somebody.
</P>

<P>
The first parchment was headed:
</P>

<P>
SCHEME FOR ECONOMY IN REALM
</P>

<P>
"Economy" caught the eye in pale pink.  The next parchment was headed:
</P>

<P>
SCHEME FOR SAFETY OF REALM
</P>

<P>
"Safety" clamoured to you in blue.
</P>

<P>
The third parchment was headed:
</P>

<P>
SCHEME FOR ENCOURAGEMENT OF LITERATURE IN REALM
</P>

<P>
"Encouragement of Literature" had got rather cramped in the small
quarters available for it.  A heading, Belvane felt, should be in one
line; she had started in letters too big for it, and the fact that the
green ink was giving out made it impossible to start afresh.
</P>

<P>
There were ten parchments altogether.
</P>

<P>
By the end of the third one, the Princess began to feel uncomfortable.
</P>

<P>
By the end of the fifth one she knew that it was a mistake her ever
having come into the Royal Family at all.
</P>

<P>
By the end of the seventh she decided that if the Countess would
forgive her this time she would never be naughty again.
</P>

<P>
By the end of the ninth one she was just going to cry.
</P>

<P>
The tenth one was in a very loud orange and was headed:
</P>

<P>
SCHEME FOR ASSISTING CALISTHENICS IN REALM
</P>

<P>
"Yes," said the Princess faintly; "I think it would be a good idea."
</P>

<P>
"I thought if your Royal Highness approved," said Belvane, "we might
just&mdash;&mdash;"
</P>

<P>
Hyacinth felt herself blushing guiltily&mdash;she couldn't think why.
</P>

<P>
"I leave it to you, Countess," she murmured.  "I am sure you know
best."
</P>

<P>
It was a remark which she would never have made to her Father.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap05"></A>
<img src="images/0089X.jpg" alt="[Illustration: Detail of Hyacinth, reviewing the Army of Amazons]">
</p>

<H3 ALIGN="center">
CHAPTER V
</H3>

<H3 ALIGN="center">
BELVANE INDULGES HER HOBBY
</H3>

<P>
In a glade in the forest the Countess Belvane was sitting: her throne,
a fallen log, her courtiers, that imaginary audience which was always
with her.  For once in her life she was nervous; she had an anxious
morning in front of her.
</P>

<P>
I can tell you the reason at once.  Her Royal Highness was going to
review her Royal Highness's Army of Amazons (see <i>Scheme II, Safety of
Realm</i>).  In half an hour she would be here.
</P>

<P>
And why not? you say.  Could anything be more gratifying?
</P>

<P>
I will tell you why not.  There was no Army of Amazons.  In order that
her Royal Highness should not know the sad truth, Belvane drew their
pay for them.  'Twas better thus.
</P>

<P>
In any trouble Belvane comforted herself by reading up her diary.  She
undid the enormous volume, and, idly turning the pages, read some of
the more delightful extracts to herself.
</P>

<P>
"<i>Monday, June 1st</i>," she read.  "Became bad."
</P>

<P>
She gave a sigh of resignation to the necessity of being bad.  Roger
Scurvilegs is of the opinion that she might have sighed a good many
years before.  According to him she was born bad.
</P>

<P>
"<i>Tuesday, June 2nd</i>," she read on.  "Realised in the privacy of my
heart that I was destined to rule the country.  <i>Wednesday, June 3rd.</i>
Decided to oust the Princess.  <i>Thursday, June 4th.</i>  Began ousting."
</P>

<P>
What a confession for any woman&mdash;even for one who had become bad last
Monday!  No wonder Belvane's diary was not for everybody.  Let us look
over her shoulder and read some more of the wicked woman's
confessions.
</P>

<P>
"<i>Friday, June 5th.</i>  Made myself a&mdash;&mdash;"  Oh, that's quite private.
However we may read this:  "<i>Thought for the week.</i>  Beware lest you
should tumble down In reaching for another's crown."  An admirable
sentiment which Roger Scurvilegs would have approved, although he
could not have rhymed it so neatly.
</P>

<P>
The Countess turned on a few more pages and prepared to write up
yesterday's events.
</P>

<P>
"<i>Tuesday, June 23rd</i>," she said to herself.  "Now what happened?
Acclaimed with enthusiasm outside the Palace&mdash;how do you spell
'enthusiasm'?"  She bit the end of her pencil and pondered.  She
turned back the pages till she came to the place.
</P>

<P>
"Yes," she said thoughtfully.  "It had three 's's' last time, so it's
'z's' turn."
</P>

<P>
She wrote "enthuzziazm" lightly in pencil; later on it would be picked
out in gold.
</P>

<P>
She closed the diary hastily.  Somebody was coming.
</P>

<P>
It was Wiggs.
</P>

<P>
"Oh, if you please, your Ladyship, her Royal Highness sent me to tell
you that she would be here at eleven o'clock to review her new army."
</P>

<P>
It was the last thing of which Belvane wanted reminding.
</P>

<P>
"Ah, Wiggs, sweet child," she said, "you find me overwhelmed."  She
gave a tragic sigh.  "Leader of the Corps de Ballet"&mdash;she indicated
with her toe how this was done, "Commander-in-Chief of the Army of
Amazons"&mdash;here she saluted, and it was certainly the least she could
do for the money, "Warden of the Antimacassars and Grand Mistress of
the Robes, I have a busy life.  Just come and dust this log for her
Royal Highness.  All this work wears me out, Wiggs, but it is my duty
and I do it."
</P>

<P>
"Woggs says you make a very good thing out of it," said Wiggs
innocently, as she began to dust.  "It must be nice to make very good
things out of things."
</P>

<P>
The Countess looked coldly at her.  It is one thing to confide to your
diary that you are bad, it's quite another to have Woggsseses shouting
it out all over the country.
</P>

<P>
"I don't know what Woggs is," said Belvane sternly, "but send it to me
at once."
</P>

<P>
As soon as Wiggs was gone, Belvane gave herself up to her passions.
She strode up and down the velvety sward, saying to herself, "Bother!
Bother!  Bother!  Bother!"  Her outbreak of violence over, she sat
gloomily down on the log and abandoned herself to despair.  Her hair
fell in two plaits down her back to her waist; on second thoughts she
arranged them in front&mdash;if one is going to despair one may as well do
it to the best advantage.
</P>

<P>
Suddenly a thought struck her.
</P>

<P>
"I am alone," she said.  "Dare I soliloquise?  I will.  It is a thing
I have not done for weeks.  'Oh, what a&mdash;&mdash;" She got up quickly.
"<i>Nobody</i> could soliloquise on a log like that," she said crossly.
She decided she could do it just as effectively when standing.  With
one pale hand raised to the skies she began again.
</P>

<P>
"Oh, what a&mdash;"
</P>

<P>
"Did you call me, Mum?" said Woggs, appearing suddenly.
</P>

<P>
"<i>Bother!</i>" said Belvane.  She gave a shrug of resignation.  "Another
time," she told herself.  She turned to Woggs.
</P>

<P>
Woggs must have been quite close at hand to have been found by Wiggs
so quickly, and I suspect her of playing in the forest when she ought
to have been doing her lessons, or mending stockings, or whatever made
up her day's work. Woggs I find nearly as difficult to explain as
Wiggs; it is a terrible thing for an author to have a lot of people
running about his book, without any invitation from him at all.
However, since Woggs is there, we must make the best of her.  I fancy
that she was a year or two younger than Wiggs and of rather inferior
education.  Witness her low innuendo about the Lady Belvane, and the
fact that she called a Countess "Mum."
</P>

<P>
"Come here," said Belvane.  "Are you what they call Woggs?"
</P>

<P>
"Please, Mum," said Woggs nervously.
</P>

<P>
The Countess winced at the "Mum," but went on bravely. "What have you
been saying about me?"
</P>

<P>
"N&mdash;Nothing, Mum."
</P>

<P>
Belvane winced again, and said, "Do you know what I do to little girls
who say things about me?  I cut their heads off; I&mdash;&mdash;"  She tried to
think of something very alarming!  "I&mdash;I stop their jam for tea.  I&mdash;I
am <i>most</i> annoyed with them."
</P>

<P>
Woggs suddenly saw what a wicked thing she had done.
</P>

<P>
"Oh, please, Mum," she said brokenly and fell on her knees.
</P>

<P>
"<i>Don't</i> call me 'Mum,'" burst out Belvane.  "It's so <i>ugly</i>.  Why do
you suppose I ever wanted to be a countess at all, Woggs, if it wasn't
so as not to be called 'Mum' any more?"
</P>

<P>
"I don't know, Mum," said Woggs.
</P>

<P>
Belvane gave it up.  The whole morning was going wrong anyhow.
</P>

<P>
"Come here, child," she sighed, "and listen.  You have been a very
naughty girl, but I'm going to let you off this time, and in return
I've something you are going to do for me."
</P>

<P>
"Yes, Mum," said Woggs.
</P>

<P>
Belvane barely shuddered now.  A sudden brilliant plan had come to
her.
</P>

<P>
"Her Royal Highness is about to review her Army of Amazons.  It is a
sudden idea of her Royal Highness's, and it comes at an unfortunate
moment, for it so happens that the Army is&mdash;er&mdash;&mdash;"  <i>What</i> was the
Army doing?  Ah, yes&mdash;"manoeuvring in a distant part of the country.
But we must not disappoint her Royal Highness.  What then shall we do,
Woggs?"
</P>

<P>
"I don't know, Mum," said Woggs stolidly.
</P>

<P>
Not having expected any real assistance from her, the Countess went
on, "I will tell you.  You see yonder tree? Armed to the teeth <i>you</i>
will march round and round it, giving the impression to one on this
side of a large army passing.  For this you will be rewarded.  Here
is&mdash;&mdash;"  She felt in the bag she carried.  "No, on second thoughts I
will owe it to you.  Now you quite understand?"
</P>

<P>
"Yes, Mum," said Woggs.
</P>

<P>
"Very well, then.  Run along to the Palace and get a sword and a
helmet and a bow and an arrow and an&mdash;an arrow and anything you like,
and then come back here and wait behind those bushes.  When I clap my
hands the army will begin to march."
</P>

<P>
Woggs curtsied and ran off.
</P>

<P>
It is probable that at this point the Countess would have resumed her
soliloquy, but we shall never know, for the next moment the Princess
and her Court were seen approaching from the other end of the glade.
Belvane advanced to meet them.
</P>

<P>
"Good morning, your Royal Highness," she said, "a beautiful day, is it
not?"
</P>

<P>
"Beautiful, Countess."
</P>

<P>
With the Court at her back, Hyacinth for the moment was less nervous
than usual, but almost at the first words of the Countess she felt her
self-confidence oozing from her.  Did I say I was like this with my
publishers?  And Roger's dragged-in Uncle&mdash;&mdash;one can't explain it.
</P>

<P>
The Court stood about in picturesque attitudes while Belvane went on:
</P>

<P>
"Your Royal Highness's brave Women Defenders, the Home Defence Army of
Amazons" (here she saluted; one soon gets into the knack of it, and it
gives an air of efficiency) "have looked forward to this day for
weeks.  How their hearts fill with pride at the thought of being
reviewed by your Royal Highness!"
</P>

<P>
She had paid, or rather received, the money for the Army so often that
she had quite got to believe in its existence. She even kept a roll of
the different companies  (it meant more delightful red ink for one
thing), and wrote herself little notes recommending Corporal Gretal
Hottshott for promotion to sergeant.
</P>

<P>
"I know very little about armies, I'm afraid," said Hyacinth.  "I've
always left that to my father.  But I think it's a sweet idea of yours
to enrol the women to defend me.  It's a little expensive, is it not?"
</P>

<P>
"Your Royal Highness, armies are <i>always</i> expensive."
</P>

<P>
The Princess took her seat, and beckoned Wiggs with a smile to her
side.  The Court, in attitudes even more picturesque than before,
grouped itself behind her.
</P>

<P>
"Is your Royal Highness ready?"
</P>

<P>
"Quite ready, Countess."
</P>

<P>
The Countess clapped her hands.
</P>

<P>
There was a moment's hesitation, and then, armed to the teeth, Amazon
after Amazon marched by. . . .
</P>

<P>
An impressive scene. . . .
</P>

<P>
However, Wiggs must needs try to spoil it.
</P>

<P>
"Why, it's Woggs!" she cried.
</P>

<P>
"Silly child!" said Belvane in an undertone, giving her a push.
</P>

<P>
The Princess looked round inquiringly.
</P>

<P>
"The absurd creature," explained the Countess, "thought she recognized
a friend in your Royal Highness's gallant Army."
</P>

<P>
"How clever of her!  They all look exactly alike to <i>me</i>."
</P>

<P>
Belvane was equal to the occasion.
</P>

<P>
"The uniform and discipline of an army have that effect rather," she
said.  "It has often been noticed."
</P>

<P>
"I suppose so," said the Princess vaguely.  "Oughtn't they to march in
fours?  I seem to remember, when I came to reviews with Father&mdash;&mdash;"
</P>

<P>
"Ah, your Royal Highness, that was an army of men.  With women&mdash;well,
we found that if they marched side by side, they <i>would</i> talk all the
time."
</P>

<P>
The Court, which had been resting on the right leg with the left knee
bent, now rested on the left leg with the right knee bent.  Woggs also
was getting tired.  The last company of the Army of Amazons was not
marching with the abandon of the first company.
</P>

<P class="noindent" align="center">
<a name="img0102"></a><img src="images/0102.jpg" alt="[Illustration: Armed to the teeth, Amazon after Amazon marched by, verso]">
<img src="images/0103.jpg" alt="[Illustration: Armed to the teeth, Amazon after Amazon marched by, recto]">
</P>

<P>
"I think I should like them to halt now so that I can address them,"
said Hyacinth.
</P>

<P>
Belvane was taken aback for the moment.
</P>

<P>
"I am afraid, your&mdash;your Royal Highness," she stammered, her brain
working busily all the time, "that that would be contrary to&mdash;to&mdash;to
the spirit of&mdash;er&mdash;the King's Regulations.  An army&mdash;an army in
marching order&mdash;must&mdash;er&mdash;<i>march</i>."  She made a long forward movement
with her hand.  "Must march," she repeated, with an innocent smile.
</P>

<P>
"I see," said Hyacinth, blushing guiltily again.
</P>

<P>
Belvane gave a loud cough.  The last veteran but two of the Army
looked inquiringly at her and passed.  The last veteran but one came
in and was greeted with a still louder cough. Rather tentatively the
last veteran of all entered and met such an unmistakable frown that it
was obvious that the march-past was over. . . .  Woggs took off her
helmet and rested in the bushes.
</P>

<P>
"That is all, your Royal Highness," said Belvane.  "158 marches past,
217 reported sick, making 622; 9 are on guard at the Palace&mdash;632 and 9
make 815.  Add 28 under age and we bring it up to the round thousand."
</P>

<P>
Wiggs opened her mouth to say something, but decided that her mistress
would probably wish to say it instead. Hyacinth, however, merely
looked unhappy.
</P>

<P>
Belvane came a little nearer.
</P>

<P>
"I&mdash;er&mdash;forgot if I mentioned to your Royal Highness that we are
paying out today.  One silver piece a day and several days in the
week, multiplied by&mdash;how many did I say?&mdash;comes to ten thousand pieces
of gold."  She produced a document, beautifully ruled.  "If your Royal
Highness would kindly initial here&mdash;&mdash;"
</P>

<P>
Mechanically the Princess signed.
</P>

<P>
"Thank you, your Royal Highness.  And now perhaps I had better go and
see about it at once."
</P>

<P>
She curtsied deeply, and then, remembering her position, saluted and
marched off.
</P>

<P>
Now Roger Scurvilegs would see her go without a pang; he would then
turn over to his next chapter, beginning "Meanwhile the King&mdash;&mdash;," and
leave you under the impression that the Countess Belvane was a common
thief.  I am no such chronicler as that.  At all costs I will be fair
to my characters.
</P>

<P>
Belvane, then, had a weakness.  She had several of which I have
already told you, but this is another one.  She had a passion for the
distribution of largesse.
</P>

<P>
I know an old gentleman who plays bowls every evening.  He trundles
his skip (or whatever he calls it) to one end of the green, toddles
after it, and trundles it back again. Think of him for a moment, and
then think of Belvane on her cream-white palfrey tossing a bag of gold
to right of her and flinging a bag of gold to left of her, as she
rides through the cheering crowds; upon my word I think hers is the
more admirable exercise.
</P>

<P>
And, I assure you, no less exacting.  When once one has got into this
habit of "flinging" or "tossing" money, to give it in any ordinary
way, to slide it gently into the palm, is unbearable.  Which of us who
has, in an heroic moment, flung half a crown to a cabman can ever be
content afterwards to hold out a handful of three-penny bits and
coppers to him?  One must always be flinging. . . .
</P>

<P>
So it was with Belvane.  The largesse habit had got hold of her.  It
is an expensive habit, but her way of doing it was less expensive than
most.  The people were taxed to pay for the Amazon Army; the pay of
the Amazon Army was flung back at them; could anything be fairer?
</P>

<P>
True, it brought her admiration and applause.  But what woman does not
like admiration?  Is that an offence?  If it is, it is something very
different from the common theft of which Roger Scurvilegs would accuse
her.  Let us be fair.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap06"></A>
<img src="images/0109X.jpg" alt="[Illustration: Detail of sleeping king]">
</p>

<H3 ALIGN="center">
CHAPTER VI
</H3>

<H3 ALIGN="center">
THERE ARE NO WIZARDS IN BARODIA
</H3>

<P>
Meanwhile "the King of Euralia was prosecuting the war with utmost
vigour."
</P>

<P>
So says Roger in that famous chapter of his, and certainly Merriwig
was very busy.
</P>

<P>
On the declaration of war the Euralian forces, in accordance with
custom, had marched into Barodia.  However hot ran the passion between
them, the two Kings always preserved the elementary courtesies of war.
The last battle had taken place in Euralian territory; this time,
therefore, Barodia was the scene of the conflict.  To Barodia, then,
King Merriwig had led his army.  Suitable pasture land had been
allotted them as a camping ground, and amid the cheers of the Barodian
populace the Euralians made their simple preparations for the night.
</P>

<P>
The two armies had now been sitting opposite to each other for some
weeks, but neither side had been idle.  On the very first morning
Merriwig had put on his Cloak of Darkness and gone to the enemy's camp
to explore the situation.  Unfortunately the same idea had occurred at
the same moment to the King of Barodia.  He also had his Cloak of
Darkness.
</P>

<P>
Half way across, to the utmost astonishment of both, the two Kings had
come violently into contact.  Realising that they had met some
unprecedented enchantment, they had hurried home after the recoil to
consult their respective Chancellors.  The Chancellors could make
nothing of it. They could only advise their Majesties to venture
another attempt on the following morning.
</P>

<P>
"But by a different route," said the Chancellors, "whereby the Magic
Pillar shall be avoided."
</P>

<P>
So by the more southerly path the two Kings ventured out next morning.
Half way across there was another violent collision, and both Kings
sat down suddenly to think it out.
</P>

<P>
"Wonder of wonders," said Merriwig.  "There is a magic wall stretching
between the two armies."
</P>

<P>
"He stood up and holding up his hand said impressively:
</P>

<P class="poem">
     "<i>Bo, boll, bill, bole.</i> <BR>
      &nbsp;<i>Wo, woll&mdash;&mdash;</i>"<BR>
</P>

<P>
"Mystery of mysteries!" cried the King of Barodia.  "It can&mdash;&mdash;"
</P>

<P>
He stopped suddenly.  Both Kings coughed.  They were remembering with
some shame their fright of yesterday.
</P>

<P>
"Who are you?" said the King of Barodia.
</P>

<P>
Merriwig saw that there was need to dissemble.
</P>

<P>
"His Majesty's swineherd," he said, in what he imagined might be a
swineherd's voice.
</P>

<P>
"Er&mdash;so am I," said the King of Barodia, rather feebly.
</P>

<P>
There was obviously nothing for it but for them to discuss swine.
</P>

<P>
Merriwig was comfortably ignorant of the subject.  The King of Barodia
knew rather less than that.
</P>

<P>
"Er&mdash;how many have you?" asked the latter.
</P>

<P>
"Seven thousand," said Merriwig at random.
</P>

<P>
"Er&mdash;so have I," said the King of Barodia, still more feebly.
</P>

<P>
"Couples," explained Merriwig.
</P>

<P>
"Mine are ones," said the King of Barodia, determined to be
independent at last.
</P>

<P>
Each King was surprised to find how easy it was to talk to an expert
on his own subject.  The King of Barodia, indeed, began to feel
reckless.
</P>

<P>
"Well," he said, "I must be getting back. It's&mdash;er&mdash;milking time."
</P>

<P>
"So must I," said Merriwig.  "By the way," he added, "what do you feed
yours on?"
</P>

<P>
The King of Barodia was not quite sure if it was apple sauce or not.
He decided that perhaps it wasn't.
</P>

<P>
"That's a secret," he said darkly.  "Been handed down from generation
to generation."
</P>

<P>
Merriwig could think of nothing better to say to this than "Ah!"  He
said it very impressively, and with a word of farewell returned to his
camp.
</P>

<P>
He was in brilliant form over the wassail bowl that night as he drew a
picture of his triumphant dissimulation.  It is only fair to say that
the King of Barodia was in brilliant form too. . . .
</P>

<P>
For several weeks after this the battle raged.  Sometimes the whole
Euralian army would line up outside its camp and call upon the
Barodians to fight; at other times the Barodian army would form fours
in full view of the Euralians in the hope of provoking a conflict.  At
intervals the two Chancellors would look up old spells, scour the
country for wizards, or send each other insulting messages.  At the
end of a month it was difficult to say which side had obtained the
advantage.
</P>

<P>
A little hill surmounted by a single tree lay half way between the two
camps.  Thither one fine morning came the two Kings and the two
Chancellors on bloody business bent.  (The phrase is Roger's.)  Their
object was nothing less than to arrange that personal fight between
the two monarchs which was always a feature of Barodo-Euralian
warfare.  The two Kings having shaken hands, their Chancellors
proceeded to settle the details.
</P>

<P>
"I suppose," said the Chancellor of Barodia, "that your Majesties will
wish to fight with swords?"
</P>

<P>
"Certainly," said the King of Barodia promptly; so promptly that
Merriwig felt certain that he had a Magic Sword too.
</P>

<P>
"Cloaks of Darkness are not allowed, of course," said the Chancellor
of Euralia.
</P>

<P>
"Why, have <i>you</i> got one?" said each King quickly to the other.
</P>

<P>
Merriwig was the first to recover himself.
</P>

<P>
"I have one&mdash;naturally," he said.  "It's a curious thing that the only
one of my subjects who has one is my&mdash;er&mdash;swineherd."
</P>

<P>
"That's funny," said the King of Barodia.  "My swineherd has one too."
</P>

<P>
"Of course," said Merriwig, "they are almost a necessity to
swineherding."
</P>

<P>
"Particularly in the milking season," said the King of Barodia.
</P>

<P>
They looked at each other with added respect.  Not many Kings in those
days had the technicalities of such a humble trade at their fingers'
ends.
</P>

<P>
The Chancellor of Barodia has been referring to the precedents.
</P>

<P>
"It was after the famous conflict between the two grandfathers of your
Majesties that the use of the Magic Cloak in personal combats was
discontinued."
</P>

<P>
"Great-grandfathers," said the Chancellor of Euralia.
</P>

<P>
"Grandfathers, I think."
</P>

<P>
"Great-grandfathers, if I am not mistaken."
</P>

<P>
Their tempers were rising rapidly, and the Chancellor of Barodia was
just about to give the Chancellor of Euralia a push when Merriwig
intervened.
</P>

<P>
"Never mind about that," he said impatiently.  "Tell us what happened
when our&mdash;our ancestors fought."
</P>

<P>
"It happened in this way, your Majesty.  Your Majesty's
grandfather&mdash;&mdash;"
</P>

<P>
"Great-grandfather," said a small voice.
</P>

<P>
The Chancellor cast one bitter look at his opponent and went on:
</P>

<P>
"The ancestors of your two Majesties arranged to settle the war of
that period by personal combat.  The two armies were drawn up in full
array.  In front of them the two monarchs shook hands.  Drawing their
swords and casting their Magic Cloaks around them, they&mdash;&mdash;"
</P>

<P>
"Well?" said Merriwig eagerly.
</P>

<P>
"It is rather a painful story, your Majesty."
</P>

<P>
"Go on, I shan't mind."
</P>

<P>
"Well, your Majesty, drawing their swords and casting their Magic
Cloaks around them they&mdash;h'r'm&mdash;returned to the wassail bowl."
</P>

<P>
"Dear, dear," said Merriwig.
</P>

<p class="noindent" align="center">
<a name="img0118"></a><img src="images/0118.jpg" alt="[Illustration: When the respective armies returned to camp they found
their Majesties asleep, verso]">
<img src="images/0119.jpg" alt="[Illustration: When the respective armies returned to camp they found
their Majesties asleep, recto]">
</P>

<P>
"When the respective armies, who had been waiting eagerly the whole of
the afternoon for some result of the combat, returned to camp, they
found their Majesties&mdash;&mdash;"
</P>

<P>
"Asleep," said the Chancellor of Euralia hastily.
</P>

<P>
"Asleep," agreed the Chancellor of Barodia.  "The excuse of their two
Majesties that they had suddenly forgotten the day, though naturally
accepted at the time, was deemed inadequate by later historians." (By
Roger and myself, anyway.)
</P>

<P>
Some further details were discussed, and then the conference closed.
The great fight was fixed for the following morning.
</P>

<P>
The day broke fine.  At an early hour Merriwig was up and practising
thrusts upon a suspended pillow.  At intervals he would consult a
little book entitled <i>Sword Play for Sovereigns</i>, and then return to
his pillow.  At breakfast he was nervous but talkative.  After
breakfast he wrote a tender letter to Hyacinth and a still more tender
one to the Countess Belvane, and burnt them.  He repeated his little
rhyme, "Bo, Boll, Bill, Bole," several times to himself until he was
word perfect.  It was just possible that it might be useful.  His last
thoughts as he rode on to the field were of his great-grandfather.
Without admiring him, he quite saw his point.
</P>

<P>
The fight was a brilliant one.  First Merriwig aimed a blow at the
King of Barodia's head which the latter parried. Then the King of
Barodia aimed a blow at his adversary's head which Merriwig parried.
This went on three or four times, and then Merriwig put into practice
a remarkable trick which the Captain of his Bodyguard had taught him.
It was his turn to parry, but instead of doing this, he struck again
at his opponent's head; and if the latter in sheer surprise had not
stumbled and fallen, there might have been a very serious ending to
the affair.
</P>

<P>
Noon found them still at it; cut and parry, cut and parry; at each
stroke the opposing armies roared their applause. When darkness put an
end to the conflict, honours were evenly divided.
</P>

<P>
It was a stiff but proud King of Euralia who received the
congratulations of his subjects that night; so proud that he had to
pour out his heart to somebody.  He wrote to his daughter.
</P>

<P>
"MY DEAR HYACINTH,
</P>

<P>
"You will be glad to hear that your father is going on well and that
Euralia is as determined as ever to uphold its honour and dignity.
To-day I fought the King of Barodia, and considering that, most
unfairly, he was using a Magic Sword, I think I may say that I did
well.  The Countess Belvane will be interested to hear that I made
4,638 strokes at my opponent and parried 4,637 strokes from him.  This
is good for a man of my age.  Do you remember that magic ointment my
aunt used to give me?  Have we any of it left?
</P>

<P>
"I played a very clever trick the other day by pretending to be a
swineherd.  I talked to a real one I met for quite a long time about
swine without his suspecting me.  The Countess might be interested to
hear this.  It would have been very awkward for me if it had been
found out who I was.
</P>

<P>
"I hope you are getting along all right.  Do you consult the Countess
Belvane at all?  I think she would be able to advise you in any
difficulties.  A young girl needs a guiding hand, and I think the
Countess would be able to advise you in any difficulties.  Do you
consult her at all?
</P>

<P>
"I am afraid this is going to be a long war.  There doesn't seem to be
a wizard in the country at all, and without one it is a little
difficult to know how to go on.  I say my spell every now and
then&mdash;you remember the one:
</P>

<P class="poem">
     '<i>Bo, boll, bill bole.</i> <BR>
      <i>Wo, woll, will, wole.</i> '<BR>
</P>

<P class="noindent">
and it certainly keeps off dragons, but we don't seem to get any
nearer defeating the enemy's army.  You might tell the Countess
Belvane that about my spell; she would be interested.
</P>

<P>
"To-morrow I go on with my fight with the King of Barodia.  I feel
quite confident now that I can hold him. He parries well, but his
cutting is not very good.  I am glad the Countess found my sword for
me; tell her that it has been most useful.
</P>

<P>
"I must now close as I must go to bed so as to be ready for my fight
to-morrow.  Good-bye, dear.  I am always,
</P>

<P align="right">
          "YOUR LOVING FATHER.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<BR>
</P>

<P>
"P.S.&mdash;I hope you are not finding your position too difficult.  If you
are in any difficulties you should consult the Countess Belvane.  I
think she would be able to advise you.  Don't forget about that
ointment.  Perhaps the Countess might know about some other kind.
It's for stiffness.  I am afraid this is going to be a long war."
</P>

<P>
The King sealed up the letter and despatched it by special messenger
the next morning.  It came to Hyacinth at a critical moment.  We shall
see in the next chapter what effect it had upon her.
</P>

<BR><BR><BR>

<P class="noindent" align="center">
<A NAME="chap07"></A>
<img src="images/0127X.jpg" alt="[Illustration: Detail of Wiggs meeting her Fairy]">
</p>

<H3 ALIGN="center">
CHAPTER VII
</H3>

<H3 ALIGN="center">
THE PRINCESS RECEIVES A LETTER AND WRITES ONE
</H3>

<P>
The Princess Hyacinth came in from her morning's ride in a very bad
temper.  She went straight up to her favourite seat on the castle
walls and sent for Wiggs.
</P>

<P>
"Wiggs," she said, "what's the matter with me?"
</P>

<P>
Wiggs looked puzzled.  She had been dusting the books in the library;
and when you dust books you simply <i>must</i> stop every now and then to
take just one little peep inside, and then you look inside another one
and another one, and by the time you have finished dusting, your head
is so full of things you have seen that you have to be asked questions
very slowly indeed.
</P>

<P>
"I'm pretty, aren't I?" went on Hyacinth.
</P>

<P>
That was an easy one.
</P>

<P>
"Lovely!" said Wiggs, with a deep breath.
</P>

<P>
"And I'm not unkind to anybody?"
</P>

<P>
"Unkind!" said Wiggs indignantly.
</P>

<P>
"Then why&mdash;oh, Wiggs, I know it's silly of me, but it <i>hurts</i> me that
my people are so much fonder of the Countess than of me."
</P>

<P>
"Oh, I'm sure they're not, your Royal Highness."
</P>

<P>
"Well, they cheer her much louder than they cheer me."
</P>

<P>
Wiggs tried to think of a way of comforting her mistress, but her head
was still full of the last book she had dusted.
</P>

<P>
"Why should they be so fond of her?" demanded Hyacinth.
</P>

<P>
"Perhaps because she's so funny," said Wiggs.
</P>

<P>
"Funny!  Is she funny?" said the Princess coldly.  "She doesn't make
<i>me</i> laugh."
</P>

<P>
"Well, it <i>was</i> funny of her to make Woggs march round and round that
tree like that, <i>wasn't</i> it?"
</P>

<P>
"Like what?  You don't mean&mdash;&mdash;"  The Princess's eyes were wide open
with astonishment.  "Was that Woggs all the time?"
</P>

<P>
"Yes, your Royal Highness.  Wasn't it lovely and funny of her?"
</P>

<P>
The Princess looked across to the forest and nodded to herself.
</P>

<P>
"Yes.  That's it.  Wiggs, I don't believe there has ever been an Army
at all. . . .  And I pay them every week!" She added solemnly, "There
are moments when I don't believe that woman is quite honest."
</P>

<P>
"Do you mean she isn't good?" asked Wiggs in awe.
</P>

<P>
Hyacinth nodded.
</P>

<P>
"I'm <i>never</i> good," said Wiggs firmly.
</P>

<P>
"What do you mean, silly?  You're the best little girl in Euralia."
</P>

<P>
"I'm <i>not</i>.  I do awful things sometimes.  Do you know what I did
yesterday?"
</P>

<P>
"Something terrible!" smiled Hyacinth.
</P>

<P>
"I tore my apron."
</P>

<P>
"You baby!  That isn't being bad," said Hyacinth absently.  She was
still thinking of that awful review.
</P>

<P>
"The Countess says it is."
</P>

<P>
"The Countess!"
</P>

<P>
"Do you know why I want to be <i>very</i> good?" said Wiggs, coming up
close to the Princess.
</P>

<P>
"Why, dear?"
</P>

<P>
"Because then I could dance like a fairy."
</P>

<P>
"Is that how it's done?" asked the Princess, rather amused.  "The
Countess must dance <i>very</i> heavily."  She suddenly remembered
something and added:  "Why, of course, child, you were going to tell
me about a fairy you met, weren't you?  That was weeks ago, though.
Tell me now.  It will help me to forget things which make me rather
angry."
</P>

<P>
It was a simple little story.  There must have been many like it in
the books which Wiggs had been dusting; but these were simple times,
and the oldest story always seemed new.
</P>

<P>
Wiggs had been by herself in the forest.  A baby rabbit had run past
her, terrified; a ferret in pursuit.  Wiggs had picked the little
fluffy thing up in her arms and comforted it; the ferret had slowed
down, walked past very indifferently with its hands, as it were, in
its pockets, hesitated a moment, and then remembered an important
letter which it had forgotten to post.  Wiggs was left alone with the
baby rabbit, and before she knew where she was, the rabbit was gone
and there was a fairy in front of her.
</P>

<P class="noindent" align="center">
<a name="img0132"></a>
<img src="images/0132.jpg"
alt="[Illustration: The rabbit was gone, and there was a fairy in front of her, verso]">
<img src="images/0133.jpg"
alt="[Illustration: The rabbit was gone, and there was a fairy in front of her, recto]">
</P>

<P>
"You have saved my life," said the fairy.  "That was a wicked magician
after me, and if he had caught me then, he would have killed me."
</P>

<P>
"Please, your Fairiness, I didn't know fairies <i>could</i> die," said
Wiggs.
</P>

<P>
"They can when they take on animal shape or human shape.  He could not
hurt me now, but before&mdash;&mdash;"  She shuddered.
</P>

<P>
"I'm so glad you're all right now," said Wiggs politely.
</P>

<P>
"Thanks to you, my child.  I must reward you.  Take this ring.  When
you have been good for a whole day, you can have one good wish; when
you have been bad for a whole day, you can have one bad wish.  One
good wish and one bad wish&mdash;that is all it will allow anybody to
have."
</P>

<P>
With these words she vanished and left Wiggs alone with the ring.
</P>

<P>
So, ever after that, Wiggs tried desperately hard to be good and have
the good wish, but it was difficult work. Something always went wrong;
she tore her apron or read books when she ought to have been dusting,
or&mdash;&mdash;  Well, you or I would probably have given it up at once, and
devoted ourselves to earning the bad wish.  But Wiggs was a nice
little girl.
</P>

<P>
"And, oh, I <i>do</i> so want to be good," said Wiggs earnestly to the
Princess, "so that I could wish to dance like a fairy."  She had a
sudden anxiety.  "That <i>is</i> a good wish, <i>isn't</i> it?"
</P>

<P>
"It's a lovely wish; but I'm sure you could dance now if you tried."
</P>

<P>
"I can't," said Wiggs.  "I always dance like this."
</P>

<P>
She jumped up and danced a few steps.  Wiggs was a dear little girl,
but her dancing reminded you of a very dusty road going up-hill all
the way, with nothing but suet-puddings waiting for you on the top.
Something like that.
</P>

<P>
"It isn't <i>really</i> graceful, is it?" she said candidly, as she came to
rest.
</P>

<P>
"Well, I suppose the fairies <i>do</i> dance better than that."
</P>

<P>
"So that's why I want to be good, so as I can have my wish."
</P>

<P>
"I really must see this ring," said the Princess.  "It sounds
fascinating."  She looked coldly in front of her and added,
"Good-morning, Countess."  (How long had the woman been there?)
</P>

<P>
"Good-morning, your Royal Highness.  I ventured to come up
unannounced.  Ah, sweet child."  She waved a caressing hand at Wiggs.
</P>

<P>
(Even if she had overheard anything, it had only been child's talk.)
</P>

<P>
"What is it?" asked the Princess.  She took a firm hold of the arms of
her chair.  She would <i>not</i>, <i>not</i>, <i>not</i> give way to the Countess
this time.
</P>

<P>
"The merest matter of business, your Royal Highness.  Just this scheme
for the Encouragement of Literature.  Your Royal Highness very wisely
decided that in the absence of the men on the sterner business of
fighting it was the part of us women to encourage the gentler arts;
and for this purpose . . . there was some talk of a competition,
and&mdash;er&mdash;&mdash;"
</P>

<P>
"Ah, yes," said Hyacinth nervously.  "I will look into that
to-morrow."
</P>

<P>
"A competition," said Belvane, gazing vaguely over Hyacinth's head.
"Some sort of a money prize," she added, as if in a trance.
</P>

<P>
"There should certainly be some sort of a prize," agreed the Princess.
(Why not, she asked herself, if one is to encourage literature?)
</P>

<P>
"Bags of gold," murmured Belvane to herself.  "Bags and bags of gold.
Big bags of silver and little bags of gold."  She saw herself tossing
them to the crowd.
</P>

<P>
"Well, we'll go into that to-morrow," said Hyacinth hastily.
</P>

<P>
"I have it all drawn up here," said Belvane.  "Your Royal Highness has
only to sign.  It saves <i>so</i> much trouble," she added with a disarming
smile. . . .  She held the document out&mdash;all in the most beautiful
colours.
</P>

<P>
Mechanically the Princess signed.
</P>

<P>
"Thank you, your Royal Highness."  She smiled again, and added, "And
now perhaps I had better see about it at once."  The Guardian of
Literature took a dignified farewell of her Sovereign and withdrew.
</P>

<P>
Hyacinth looked at Wiggs in despair.
</P>

<P>
"There!" she said.  "That's me.  I don't know what it is about that
woman, but I feel just a child in front of her.  Oh, Wiggs, Wiggs, I
feel so lonely sometimes with nothing but women all around me.  I wish
I had a man here to help me."
</P>

<P>
"Are <i>all</i> the men fighting in <i>all</i> the countries?"
</P>

<P>
"Not all the countries.  There's&mdash;Araby.  Don't you remember&mdash;oh, but
of course you wouldn't know anything about it.  But Father was just
going to ask Prince Udo of Araby to come here on a visit, when the war
broke out.  Oh, I wish, I <i>wish</i> Father were back again."  She laid
her head on her arms; and whether she would have shed a few royal
tears or had a good homely cry, I cannot tell you.  For at that moment
an attendant came in.  Hyacinth was herself again at once.
</P>

<P>
"There is a messenger approaching on a horse, your Royal Highness,"
she announced.  "Doubtless from His Majesty's camp."
</P>

<P>
With a shriek of delight, and an entire lack of royal dignity, the
Princess, followed by the faithful Wiggs, rushed down to receive him.
</P>

<P>
Meanwhile, what of the Countess?  She was still in the Palace, and,
more than that, she was in the Throne Room of the Palace, and, more
even than that, she was on the Throne, of the Throne Room of the
Palace.
</P>

<P>
She couldn't resist it.  The door was open as she came down from her
interview with the Princess, and she had to go in.  There was a woman
in there, tidying up, who looked questioningly at Belvane as she
entered.
</P>

<P>
"You may leave," said the Countess with dignity.  "Her Royal Highness
sent me in here to wait for her."
</P>

<P>
The woman curtsied and withdrew.
</P>

<P>
The Countess then uttered these extraordinary words:
</P>

<P>
"When I am Queen in Euralia they shall leave me backwards!"
</P>

<P>
Her subsequent behaviour was even more amazing.
</P>

<P>
She stood by the side of the door, and putting her hand to her mouth
said shrilly, "Ter-rum, ter-rum, terrumty-umty-um."  Then she took her
hand away and announced loudly, "Her Majesty Queen Belvane the First!"
after which she cheered slightly.
</P>

<P>
Then in came Her Majesty, a very proper dignified gracious Queen&mdash;none
of your seventeen-year-old chits.  Bowing condescendingly from side to
side she made her way to the Throne, and with a sweep of her train she
sat down.
</P>

<P>
Courtiers were presented to her; representatives from foreign
countries; Prince Hanspatch of Tregong, Prince Ulric, the Duke of
Highanlow.
</P>

<P>
"Ah, my dear Prince Hanspatch," she cried, stretching out her hand to
the right of her; "and you, dear Prince Ulric," with a graceful
movement of the left arm towards him; "and, dear Duke, <i>you</i> also!"
Her right hand, which Prince Hanspatch had by now finished with, went
out to the Duke of Highanlow that he too might kiss it.
</P>

<P>
But it was arrested in mid-air.  She felt rather than saw that the
Princess was watching her in amazement from the doorway.
</P>

<P>
Without looking round she stretched out again first one arm and then
the other.  Then, as if she had just seen the Princess, she jumped up
in a pretty confusion.
</P>

<P>
"Oh, your Royal Highness," she cried, "you caught me at my physical
exercises!"  She gave a self-conscious little laugh.  "My physical
exercises&mdash;a forearm movement."  Once again she stretched out her arm.
"Building up the&mdash;er&mdash;building up&mdash;building up&mdash;&mdash;"
</P>

<P>
Her voice died away, for the Princess still looked coldly at her.
</P>

<P>
"Charming, Countess," she said.  "I am sorry to interrupt you, but I
have some news for you.  You will like to know that I am inviting
Prince Udo of Araby here on a visit.  I feel we want a little outside
help in our affairs."
</P>

<P>
"Prince Udo?" cried the Countess.  "<i>Here?</i>"
</P>

<P>
"Have you any objection?" said Hyacinth.  She found it easier to be
stern now, for the invitation had already been sent off by the hand of
the King's Messenger.  Nothing that the Countess could say could
influence her.
</P>

<P>
"No objection, your Royal Highness; but it seems so strange.  And then
the expense!  Men are such hearty eaters.  Besides," she looked with a
charming smile from the Princess to Wiggs, "we were all getting on so
<i>nicely</i> together!  Of course if he just dropped in for afternoon tea
one day&mdash;&mdash;"
</P>

<P>
"He will make a stay of some months, I hope."  There were no wizards
in Barodia, and therefore the war would be a long one.  It was this
which had decided Hyacinth.
</P>

<P>
"Of course," said Belvane, "whatever your Royal Highness wishes, but I
do think that His Majesty&mdash;&mdash;"
</P>

<P>
"My dear Countess," said Hyacinth, with a smile, "the invitation has
already gone, so there's nothing more to be said, is there?  Had you
finished your exercises?  Yes?  Then, Wiggs, will you conduct her
ladyship downstairs?"
</P>

<P>
She turned and left her.  The Countess watched her go, and then stood
tragically in the middle of the room, clasping her diary to her
breast.
</P>

<P>
"This is terrible!" she said.  "I feel <i>years</i> older."  She held out
her diary at arm's length and said in a gloomy voice, "<i>What</i> an entry
for to-morrow!"  The thought cheered her up a little.  She began to
consider plans.  How could she circumvent this terrible young man who
was going to put them all in their places.  She wished that&mdash;&mdash;
</P>

<P>
All at once she remembered something.
</P>

<P>
"Wiggs," she said, "what was it I heard you saying to the Princess
about a wish?"
</P>

<P>
"Oh, that's my ring," said Wiggs eagerly.  "If you've been good for a
whole day you can have a good wish.  And my wish is that&mdash;&mdash;"
</P>

<P>
"A wish!" said Belvane to herself.  "Well, I wish that&mdash;&mdash;"  A sudden
thought struck her.  "You said that you had to be good for a whole day
first?"
</P>

<P>
"Yes."
</P>

<P>
Belvane mused.
</P>

<P>
"I wonder what they mean by <i>good</i>," she said.
</P>

<P>
"Of course," explained Wiggs, "if you've been bad for a whole day you
can have a bad wish.  But I should hate to have a bad wish, wouldn't
you?"
</P>

<P>
"Simply hate it, child," said Belvane.  "Er&mdash;may I have a look at that
ring?"
</P>

<P>
"Here it is," said Wiggs; "I always wear it round my neck."
</P>

<P>
The Countess took it from her.
</P>

<P>
"Listen," she said.  "Wasn't that the Princess calling you?  Run
along, quickly, child."  She almost pushed her from the room and
closed the door on her.
</P>

<P>
Alone again, she paced from end to end of the great chamber, her left
hand nursing her right elbow, her chin in her right hand.
</P>

<P>
"If you are good for a day," she mused, "you can have a good wish.  If
you are bad for a day you can have a bad wish.  Yesterday I drew ten
thousand pieces of gold for the Army; the actual expenses were what I
paid&mdash;what I owe Woggs. . . .  I suppose that is what narrow-minded
people call being bad. . . .  I suppose this Prince Udo would call it
bad. . . .  I suppose he thinks he will marry the Princess and throw
me into prison."  She flung her head back proudly.  "Never!"
</P>

<P>
Standing in the middle of the great Throne Room, she held the ring up
in her two hands and wished.
</P>

<P>
"I wish," she said, and there was a terrible smile in her eyes, "I
wish that something very&mdash;very <i>humorous</i> shall happen to Prince Udo
on his journey."
</P>

<BR><BR><BR>

<P class="noindent" align="center">
<A NAME="chap08"></A>
<img src="images/0147X.jpg" alt="[Illustration: Detail of Udo and Coronel on their journey]">
</P>

<H3 ALIGN="center">
CHAPTER VIII
</H3>

<H3 ALIGN="center">
PRINCE UDO SLEEPS BADLY
</H3>

<P>
Everybody likes to make a good impression on his first visit, but
there were moments just before his arrival in Euralia when Prince Udo
doubted whether the affair would go as well as he had hoped.  You
shall hear why.
</P>

<P>
He had been out hunting with his friend, the young Duke Coronel, and
was returning to the Palace when Hyacinth's messenger met him.  He
took the letter from him, broke the seals, and unrolled it.
</P>

<P>
"Wait a moment, Coronel," he said to his friend.  "This is going to be
an adventure of some sort, and if it's an adventure I shall want you
with me."
</P>

<P>
"I'm in no hurry," said Coronel, and he got off his horse and gave it
into the care of an attendant.  The road crossed a stream here.
Coronel sat up on the little stone bridge and dropped pebbles idly
into the water.
</P>

<P>
The Prince read his letter.
</P>

<P>
<i>Plop . . . Plop . . . Plop . . . Plop . . .</i>
</P>

<P>
The Prince looked up from his letter.
</P>

<P>
"How many days' journey is it to Euralia?" he asked Coronel.
</P>

<P>
"How long did it take the messenger to come?" answered Coronel,
without looking up.  (<i>Plop.</i> )
</P>

<P>
"I might have thought of that myself," said Udo, "only this letter has
rather upset me."  He turned to the messenger. "How long has it&mdash;&mdash;?"
</P>

<P>
"Isn't the letter dated?" said Coronel.  (<i>Plop.</i> )
</P>

<P>
Udo paid no attention to this interruption and finished his question
to the messenger.
</P>

<P>
"A week, sire."
</P>

<P>
"Ride on to the castle and wait for me.  I shall have a message for
you."
</P>

<P>
"What is it?" said Coronel, when the messenger had gone.  "An
adventure?"
</P>

<P>
"I think so.  I think we may call it that, Coronel."
</P>

<P>
"With me in it?"
</P>

<P>
"Yes, I think you will be somewhere in it."
</P>

<P>
Coronel stopped dropping his pebbles and turned to the Prince.
</P>

<P>
"May I hear about it?"
</P>

<P>
Udo help out the letter; then feeling that a lady's letter should be
private, drew it back again.  He prided himself always on doing the
correct thing.
</P>

<P>
"It's from Princess Hyacinth of Euralia," he said; "she doesn't say
much.  Her father is away fighting, and she is alone and she is in
some trouble or other.  It ought to make rather a good adventure."
</P>

<P>
Coronel turned away and began to drop his pebbles into the stream
again.
</P>

<P>
"Well, I wish you luck," he said.  "If it's a dragon, don't forget
that&mdash;&mdash;"
</P>

<P>
"But you're coming, too," said Udo, in dismay.  "I must have you with
me."
</P>

<P>
"Doing what?"
</P>

<P>
"What?"
</P>

<P>
"Doing what?" said Coronel again.
</P>

<P>
"Well," said Prince Udo awkwardly, "er&mdash;well, you&mdash;well."
</P>

<P>
He felt that it was a silly question for Coronel to have asked.
Coronel knew perfectly well what he would be doing all the time.  In
Udo's absence he would be telling Princess Hyacinth stories of his
Royal Highness's matchless courage and wisdom.  An occasional
discussion also with the Princess upon the types of masculine beauty,
leading up to casual mention of Prince Udo's own appearance, would be
quite in order.  When Prince Udo was present Coronel would no doubt
find the opportunity of drawing Prince Udo out, an opportunity of
which a stranger could not so readily avail himself.
</P>

<P>
But of course you couldn't very well tell Coronel that.  A man of any
tact would have seen it at once.
</P>

<P>
"Of course," he said, "don't come if you don't like.  But it would
look rather funny if I went quite unattended; and&mdash;and her Royal
Highness is said to be very beautiful," he added lamely.
</P>

<P>
Coronel laughed.  There are adventures and adventures; to sit next to
a very beautiful Princess and discuss with her the good looks of
another man was not the sort of adventure that Coronel was looking
for.
</P>

<P>
He tossed the remainder of his pebbles into the stream and stood up.
</P>

<P>
"Of course, if your Royal Highness wishes&mdash;&mdash;"
</P>

<P>
"Don't be a fool, Coronel," said his Royal Highness, rather snappily.
</P>

<P>
"Well, then, I'll come with my good friend Udo if he wants me."
</P>

<P>
"I do want you."
</P>

<P>
"Very well, that settles it.  After all," he added to himself, "there
may be <i>two</i> dragons."
</P>

<P>
Two dragons would be one each.  But from all accounts there were not
two Princesses.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
So three days later the friends set out with good hearts upon the
adventure.  The messenger had been sent back to announce their
arrival; they gave him three days' start, and hoped to gain two days
upon him.  In the simple fashion of those times (so it would seem from
Roger Scurvilegs) they set out with no luggage and no clear idea of
where they were going to sleep at night.  This, after all, is the best
spirit in which to start a journey.  It is the Gladstone bag which has
killed romance.
</P>

<P>
They started on a perfect summer day, and they rode past towers and
battlements, and by the side of sparkling streams, and came out into
the sunlight again above sleepy villages, and, as they rode, Coronel
sang aloud and Udo tossed his sword into the air and caught it again.
As evening fell they came to a woodman's cottage at the foot of a high
hill, and there they decided to rest for the night.  An old woman came
out to welcome them.
</P>

<P>
"Good evening, your Royal Highness," she said.
</P>

<P class="noindent" align="center">
<a name="img0154"></a>
<img src="images/0154.jpg" alt="[Illustration: As evening fell they came to a woodman's cottage at
the foot of a high hill, verso]">
<img src="images/0155.jpg" alt="[Illustration: As evening fell they came to a woodman's cottage at the foot of a high hill, recto]">
</P>

<P>
"You know me?" said Udo, more pleased than surprised.
</P>

<P>
"I know all who come into my house," said the old woman solemnly, "and
all who go away from it."
</P>

<P>
This sort of conversation made Coronel feel creepy.  There seemed to
be a distinction between the people who came to the house and the
people who went away from it which he did not like.
</P>

<P>
"Can we stay here the night, my good woman?" said Udo.
</P>

<P>
"You have hurt your hand," she said, taking no notice of his question.
</P>

<P>
"It's nothing," said Udo hastily.  On one occasion he had caught his
sword by the sharp end by mistake&mdash;a foolish thing to have done.
</P>

<P>
"Ah, well, since you won't want hands where you're going, it won't
matter much."
</P>

<P>
It was the sort of thing old women said in those days, and Udo did not
pay much attention to it.
</P>

<P>
"Yes, yes," he said; "but can you give my friend and myself a bed for
to-night?"
</P>

<P>
"Seeing that you won't be travelling together long, come in and
welcome."
</P>

<P>
She opened the door and they followed her in.
</P>

<P>
As they crossed the threshold, Udo half turned round and whispered
over his shoulder to Coronel,
</P>

<P>
"Probably a fairy.  Be kind to her."
</P>

<P>
"How can one be kind to one's hostess?" said Coronel.  "It's she who
has to be kind to <i>us</i>."
</P>

<P>
"Well, you know what I mean; don't be rude to her."
</P>

<P>
"My dear Udo, this to <i>me</i>&mdash;the pride of Araby, the favourite courtier
of his Majesty, the&mdash;&mdash;"
</P>

<P>
"Oh, all right," said Udo.
</P>

<P>
"Sit down and rest yourselves," said the old woman.  "There'll be
something in the pot for you directly."
</P>

<P>
"Good," said Udo.  He looked approvingly at the large cauldron hanging
over the fire.  It was a big fireplace for such a small room.  So he
thought when he first looked at it, but as he gazed, the room seemed
to get bigger and bigger, and the fireplace to get farther and farther
away, until he felt that he was in a vast cavern cut deep into the
mountainside.  He rubbed his eyes, and there he was in the small
kitchen again and the cauldron was sending out a savoury smell.
</P>

<P>
"There'll be something in it for all tastes," went on the old woman,
"even for Prince Udo's."
</P>

<P>
"I'm not so particular as all that," said Udo mildly.  The room had
just become five hundred yards long again, and he was feeling quiet.
</P>

<P>
"Not now, but you will be."
</P>

<P>
She filled them a plate each from the pot; and pulling their chairs up
to the table, they fell to heartily.
</P>

<P>
"This is really excellent," said Udo, as he put down his spoon and
rested for a moment.
</P>

<P>
"You'd think you'd always like that, wouldn't you?" she said.
</P>

<P>
"I always shall be fond of anything so perfectly cooked."
</P>

<P>
"Ah," remarked the old woman thoughtfully.
</P>

<P>
Udo was beginning to dislike her particular style of conversation.  It
seemed to carry the merest suggestion of a hint that something
unpleasant was going to happen to him.  Nothing apparently was going
to happen to Coronel.  He tried to drag Coronel into the conversation
in case the old woman had anything over for him.
</P>

<P>
"My friend and I," he said, "hope to be in Euralia the day after
to-morrow."
</P>

<P>
"No harm in hoping," was the answer.
</P>

<P>
"Dear me, is something going to happen to us on the way?"
</P>

<P>
"Depends what you call 'us.'"
</P>

<P>
Coronel pushed back his chair and got up.
</P>

<P>
"I know what's going to happen to me," he said.  "I'm going to sleep."
</P>

<P>
"Well," said Udo, getting up too, "we've got a long day before us
to-morrow, and apparently we are in for an adventure&mdash;er, <i>we</i> are in
for an adventure of some sort."  He looked anxiously at the old woman,
but she made no sign.  "And so let's to bed."
</P>

<P>
"This way," said the old woman, and by the light of a candle she led
them upstairs.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
Udo slept badly.  He had a feeling (just as you have) that something
was going to happen to him; and it was with some surprise that he woke
up in the morning to find himself much as he was when he went to bed.
He looked at himself in the glass; he invited Coronel to gaze at him;
but neither could discover that anything was the matter.
</P>

<P>
"After all," said Udo, "I don't suppose she meant anything.  These old
women get into a way of talking like that.  If anybody is going to be
turned into anything, it's much more likely to be you."
</P>

<P>
"Is that why you brought me with you?" asked Coronel.
</P>

<P>
I suppose that by this time they had finished their dressing.  Roger
Scurvilegs tells us nothing on such important matters; no doubt from
modesty.  "Next morning they rose," he says, and disappoints us of a
picture of Udo brushing his hair.  They rose and went down to
breakfast.
</P>

<P>
The old woman was in a less cryptic mood at breakfast.  She was
particularly hospitable to Udo, and from some secret store produced an
unending variety of good things for him to eat.  To Coronel it almost
looked as if she were fattening him up for something, but this
suggestion was received with such bad grace by Udo that he did not
pursue the subject.
</P>

<P>
As soon as breakfast was over they started off again.  From one of the
many bags of gold he carried, Udo had offered some acknowledgment to
the old woman, but she had refused to take it.
</P>

<P>
"Nay, nay," she said.  "I shall be amply rewarded before the day is
out."  And she seemed to be smiling to herself as if she knew of some
joke which the Prince and Coronel did not yet share.
</P>

<P>
"I like to-day," said Coronel as they rode along.  "There's a smell of
adventure in the air.  Red roofs, green trees, blue sky, white road&mdash;I
could fall in love to-day."
</P>

<P>
"Who with?" said Udo suspiciously.
</P>

<P>
"Any one&mdash;that old woman, if you like."
</P>

<P>
"Oh, don't talk of her," said the Prince with a shudder. "Coronel,
hadn't you a sense of being <i>out</i> of some joke that she was in?"
</P>

<P>
"Perhaps we shall be in it before long.  I could laugh very easily on
a morning like this."
</P>

<P>
"Oh, I can see a joke as well as any one," said Udo.  "Don't be afraid
that I shan't laugh, too.  No doubt it will make a good story,
whatever it is, to tell to the Princess Hyacinth.  Coronel," he added
solemnly, the thought having evidently only just occurred to him,  "I
am all impatience to help that poor girl in her trouble."  And as if
to show his impatience, he suddenly gave the reins a shake and
cantered ahead of his companion.  Smiling to himself, Coronel followed
at his leisure.
</P>

<P>
They halted at mid-day in a wood, and made a meal from some provisions
which the old woman had given them; and after they had eaten, Udo lay
down on a mossy bank and closed his eyes.
</P>

<P>
"I'm sleepy," he said; "I had a restless night.  Let's stay here
awhile; after all, there's no hurry."
</P>

<P>
"Personally," said Coronel, "I'm all impatience to help that&mdash;&mdash;"
</P>

<P>
"I tell you I had a very bad night," said Udo crossly.
</P>

<P>
"Oh, well, I shall go off and look for dragons.  Coronel, the Dragon
Slayer.  Good-bye."
</P>

<P>
"Only half an hour," said Udo.
</P>

<P>
"Right."
</P>

<P>
With a nod to the Prince he strolled off among the trees.
</P>

<BR>

<P align="right"><img src="images/0164X.jpg" alt="[Illustration: Small decoration of Belvane writing in her diary.]">
</P>

<BR><BR><BR>

<P class="noindent" align="center">
<A NAME="chap09"></A>
<img src="images/0165X.jpg" alt="[Illustration: Detail of Udo in his animal form, coming out of some plants.]">
</P>

<H3 ALIGN="center">
CHAPTER IX
</H3>

<H3 ALIGN="center">
THEY ARE AFRAID OF UDO
</H3>

<P>
This is a painful chapter for me to write.  Mercifully it is to be a
short one.  Later on I shall become used to the situation; inclined,
even, to dwell upon its humorous side; but for the moment I cannot see
beyond the sadness of it.  That to a Prince of the Royal House of
Araby, and such an estimable young man as Udo, those things should
happen. Roger Scurvilegs frankly breaks down over it.  "That
abominable woman," he says (meaning, of course, Belvane), and he has
hysterics for more than a page.
</P>

<P>
Let us describe it calmly.
</P>

<P>
Coronel came back from his stroll in the same casual way in which he
had started and dropped down lazily upon the grass to wait until Udo
was ready to mount.  He was not thinking of Udo.  He was wondering if
Princess Hyacinth had an attendant of surpassing beauty, or a dragon
of surpassing malevolence&mdash;if, in fact, there were any adventures in
Euralia for a humble fellow like himself.
</P>

<P>
"Coronel!" said a small voice behind him.
</P>

<P>
He turned round indifferently.
</P>

<P>
"Hullo, Udo, where are you?" he said.  "Isn't it time we were
starting?"
</P>

<P>
"We aren't starting," said the voice.
</P>

<P>
"What's the matter?  What are you hiding in the bushes for?
Whatever's the matter, Udo?"
</P>

<P>
"I'm not very well."
</P>

<P>
"My poor Udo, what's happened?"  He jumped up and made towards him.
</P>

<P>
"Stop!" shrieked the voice.  "I command you!"
</P>

<P>
Coronel stopped.
</P>

<P>
"Your Royal Highness's commands," he began rather coldly&mdash;&mdash;
</P>

<P>
There was an ominous sniffing from the bushes.
</P>

<P>
"Coronel," said an unhappy voice at last, "I think I'm coming out."
</P>

<P>
Wondering what it all meant, Coronel waited in silence.
</P>

<P>
"Yes, I am coming out, Coronel," said the voice.  "But you mustn't be
surprised if I don't look very well.  I'm&mdash;I'm&mdash;Coronel, here I am,"
said Udo pathetically and he stepped out.
</P>

<P>
Coronel didn't know whether to laugh or to cry.
</P>

<P>
Poor Prince Udo!
</P>

<P class="noindent" align="center">
<a name="img0168"></a><img src="images/0168.jpg" alt="[Illustration: &quot;Coronel, here I am,&quot; said Udo pathetically, and he stepped out, verso]">
<img src="images/0169.jpg" alt="[Illustration: &quot;Coronel, here I am,&quot; said Udo pathetically, and he stepped out, recto]">
</P>

<P>
He had the head and the long ears of a rabbit, and in some unfortunate
way a look of the real Prince Udo in spite of it.  He had the mane and
the tail of a lion.  In between the tail and the mane it is difficult
to say what he was, save that there was an impression of magnificence
about his person&mdash;such magnificence, anyhow, as is given by an
astrakhan-trimmed fur coat.
</P>

<P>
Coronel decided that it was an occasion for tact.
</P>

<P>
"Ah, here you are," he said cheerfully.  "Shall we get along?"
</P>

<P>
"Don't be a fool, Coronel," said Udo, almost crying.  "Don't pretend
that you can't <i>see</i> that I've got a tail."
</P>

<P>
"Why, bless my soul, so you have.  A tail!  Well, think of that!"
</P>

<P>
Udo showed what he thought of it by waving it peevishly.
</P>

<P>
"This is not a time for tact," he said.  "Tell me what I look like."
</P>

<P>
Coronel considered for a moment.
</P>

<P>
"Really frankly?" he asked.
</P>

<P>
"Y&mdash;yes," said Udo nervously.
</P>

<P>
"Then, frankly, your Royal Highness looks&mdash;funny."
</P>

<P>
"<i>Very</i> funny?" said Udo wistfully.
</P>

<P>
"<i>Very</i> funny," said Coronel.
</P>

<P>
His Highness sighed.
</P>

<P>
"I was afraid so," he said.  "That's the cruel part about it.  Had I
been a lion there would have been a certain pathetic splendour about
my position.  Isolated&mdash;cut off&mdash;suffering in regal silence."  He
waved an explanatory paw.  "Even in the most hideous of beasts there
might be a dignity."  He meditated for a moment.  "Have you ever seen
a yak, Coronel?" he asked.
</P>

<P>
"Never."
</P>

<P>
"I saw one once in Barodia.  It is not a beautiful animal, Coronel;
but as a yak I should not have been entirely unlovable.  One does not
laugh at a yak, Coronel, and where one does not laugh one may come to
love. . . .  What does my head look like?"
</P>

<P>
"It looks&mdash;striking."
</P>

<P>
"I haven't seen it, you see."
</P>

<P>
"To one who didn't know your Royal Highness it would convey the
impression of a rabbit."
</P>

<P>
Udo laid his head between his paws and wept.
</P>

<P>
"A r&mdash;rabbit!" he sobbed.  So undignified, so lacking in true pathos,
so&mdash;&mdash;  And not even a whole rabbit," he added bitterly.
</P>

<P>
"How did it happen?"
</P>

<P>
"I don't know, Coronel.  I just went to sleep, and woke up feeling
rather funny, and&mdash;&mdash;"  He sat up suddenly and stared at Coronel.  "It
was that old woman did it.  You mark my words, Coronel; she did it."
</P>

<P>
"Why should she?"
</P>

<P>
"I don't know.  I was very polite to her.  Don't you remember my
saying to you, 'Be polite to her, because she's probably a fairy!'
You see, I saw through her disguise at once.  Coronel, what shall we
do?  Let's hold a council of war and think it over."
</P>

<P>
So they held a council of war.
</P>

<P>
Prince Udo put forward two suggestions.
</P>

<P>
The first was that Coronel should go back on the morrow and kill the
old woman.
</P>

<P>
The second was that Coronel should go back that afternoon and kill the
old woman.
</P>

<P>
Coronel pointed out that as she had turned Prince Udo into&mdash;into
a&mdash;a&mdash;("Quite so," said Udo)&mdash;it was likely that she alone could turn
him back again, and that in that case he had better only threaten her.
</P>

<P>
"I want <i>somebody</i> killed," said Udo, rather naturally.
</P>

<P>
"Suppose," said Coronel, "you stay here for two days while I go back
and see the old witch, and make her tell me what she knows.  She knows
something, I'm certain.  Then we shall see better what to do."
</P>

<P>
Udo mused for a space.
</P>

<P>
"Why didn't they turn <i>you</i> into anything?" he asked.
</P>

<P>
"Really, I don't know.  Perhaps because I'm too unimportant."
</P>

<P>
"Yes, that must be it."  He began to feel a little brighter.
"Obviously, that's it."  He caressed a whisker with one of his paws.
"They were afraid of me."
</P>

<P>
He began to look so much happier that Coronel thought it was a
favourable moment in which to withdraw.
</P>

<P>
"Shall I go now, your Royal Highness?"
</P>

<P>
"Yes, yes, you may leave me."
</P>

<P>
"And shall I find you here when I come back?"
</P>

<P>
"You may or you may not, Coronel; you may or you may not. . . .
Afraid of me," he murmured to himself.  "Obviously."
</P>

<P>
"And if I don't?"
</P>

<P>
"Then return to the Palace."
</P>

<P>
"Good-bye, your Royal Highness."
</P>

<P>
Udo waved a paw at him.
</P>

<P>
"Good-bye, good-bye."
</P>

<P>
Coronel got on his horse and rode away.  As soon as he was out of
earshot he began to laugh.  Spasm after spasm shook him.  No sooner
had he composed himself to gravity than a remembrance of Udo's
appearance started him off again.
</P>

<P>
"I couldn't have stayed with him a moment longer," he thought.  "I
should have burst.  Poor Udo!  However, we'll soon get him all right."
</P>

<P>
That evening he reached the place where the cottage had stood, but it
was gone.  Next morning he rode back to the wood.  Udo was gone too.
He returned to the Palace, and began to think it out.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
Left to himself Udo very soon made up his mind.  There were three
courses open to him.
</P>

<P>
He might stay where he was till he was restored to health.
</P>

<P>
This he rejected at once.  When you have the head of a rabbit, the
tail of a lion, and the middle of a woolly lamb, the need for action
of some kind is imperative.  All the blood of your diverse ancestors
calls to you to be up and doing.
</P>

<P>
He might go back to Araby.
</P>

<P>
To Araby, where he was so well-known, so respected, so popular?  To
Araby, where he rode daily among his father's subjects that they might
have the pleasure of cheering him?  How awkward for everybody!
</P>

<P>
On to Euralia then?
</P>

<P>
Why not?  The Princess Hyacinth had called for him.  What devotion it
showed if he came to her even now&mdash;in his present state of bad health!
She was in trouble: enchanters, wizards, what-nots.  Already, then,
he had suffered in her service&mdash;so at least he would say, and so
possibly it might be.  Coronel had thought him&mdash;funny; but women had
not much sense of humour as a rule.  Probably as a child Hyacinth had
kept rabbits . . . or lambs.  She would find him&mdash;strokable. . . .
And the lion in him . . . in his tail, his fierce mane . . . she would
find that inspiring.  Women like to feel that there is something
fierce, untamable in the man they love; well, there it was.
</P>

<P>
It was not as if he had Coronel with him.  Coronel and he (in his
present health) could never have gone into Euralia together; the
contrast was too striking; but he alone, Hyacinth's only help!  Surely
she would appreciate his magnanimity.
</P>

<P>
Also, as he had told himself a moment ago, there was quite a chance
that it was a Euralian enchanter who had put this upon him&mdash;to prevent
him helping Hyacinth.  If so, he had better go to Euralia in order to
deal with that enchanter.  For the moment, he did not see exactly how
to deal with him, but no doubt he would think of some tremendously
cunning device later on.
</P>

<P>
To Euralia then with all dispatch.
</P>

<P>
He trotted off.  As Coronel had said, they were evidently afraid of
him.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap10"></A><img src="images/0179X.jpg" alt="[Illustration: Detail of Belvane on horseback and throwing something]">
</p>


<H3 ALIGN="center">
CHAPTER X
</H3>

<H3 ALIGN="center">
CHARLOTTE PATACAKE ASTONISHES THE CRITICS
</H3>

<P>
The Lady Belvane sits in her garden.  She is very happy. An enormous
quill-pen, taken from a former favourite goose and coloured red, is in
her right hand.  The hair of her dark head, held on one side, touches
the paper whereon she writes, and her little tongue peeps out between
her red lips.  Her left hand taps the table&mdash;one-two, one-two,
one-two, one-two, one-two.  She is composing.
</P>

<P>
Wonderful woman!
</P>

<P>
You remember that scene with the Princess Hyacinth?  "I feel we want a
little outside help in our affairs."  A fortnight of suspense before
Prince Udo arrived.  What had the ring done to him?  At the best, even
if there would be no Udo at all to interfere, nevertheless she knew
that she had lost her footing at the Palace.  She and the Princess
would now be open enemies.  At the worst&mdash;those magic rings were so
untrustworthy!&mdash;a Prince, still powerful, and now seriously annoyed,
might be leagued against her.
</P>

<P>
Yet she composed.
</P>

<P>
And what is she writing?  She is entering for the competition in
connection with the Encouragement of Literature Scheme: the last
scheme which the Princess had signed.
</P>

<P>
I like to think of her peacefully writing at a time when her whole
future hung in the balance.  Roger sneers at her.  "Even now," he
says, "she was hoping to wring a last bag-full of gold from her
wretched country."  I deny emphatically that she was doing anything of
the sort.  She was entering for a duly authorised competition under
the pen-name of Charlotte Patacake.  The fact that the Countess
Belvane, according to the provisions of the scheme, was sole judge of
the competition, is beside the point.  Belvane's opinion of Charlotte
Patacake's poetry was utterly sincere, and uninfluenced in any way by
monetary considerations.  If Patacake were rewarded the first prize it
would be because Belvane honestly thought she was worth it.
</P>

<P>
One other fact by way of defence against Roger's slanders.  As judge,
Belvane had chosen the subject of the prize poems.  Now Belvane and
Patacake both excelled in the lighter forms of lyrical verse; yet the
subject of the poem was to be epic.  "The Barodo-Euralian War"&mdash;no
less.  How many modern writers would be as fair?
</P>

<P class="poem">
"THE BARODO-EURALIAN WAR."
</P>

<P>
This line is written in gold, and by itself would obtain a prize in
any local competition.
</P>

<P class="poem">
    <i>King Merriwig the First rode out to war</i><BR>
    <i>As many other kings had done before!</i><BR>
    <i>Five hundred men behind him marched to fight&mdash;</i><BR>
</P>

<P>
There follows a good deal of scratching out, and then comes (a sudden
inspiration) this sublimely simple line:
</P>

<P class="poem">
    <i>Left-right, left-right, left-right, left-right, left-right.</i> <BR>
</P>

<P>
One can almost hear the men moving.
</P>

<P class="poem">
    <i>What gladsome cheers assailed the balmy air&mdash;</i><BR>
    <i>They came from north, from south, from everywhere!</i><BR>
    <i>No wight that stood upon that sacred scene</i><BR>
    <i>Could gaze upon the sight unmoved, I ween:</i><BR>
    <i>No wight that stood upon that sacred spot</i><BR>
    <i>Could gaze upon the sight unmoved, I wot:</i><BR>
</P>

<P>
It is not quite clear whether the last couplet is an alternative to
the couplet before or is purposely added in order to strengthen it.
Looking over her left shoulder it seems to me that there is a line
drawn through the first one, but I cannot see very clearly because of
her hair, which will keep straying over the page.
</P>

<P class="poem">
    <i>Why do they march so fearless and so bold?</i><BR>
    <i>The answer is not very quickly told.</i> <BR>
    <i>To put it shortly, the Barodian king</i><BR>
    <i>Insulted Merriwig like anything&mdash;</i><BR>
    <i>King Merriwig, the dignified and wise,</i> <BR>
    <i>Who saw him flying over with surprise,</i> <BR>
    <i>As did his daughter, Princess Hyacinth.</i> <BR>
</P>

<P>
This was as far as she had got.
</P>

<P>
She left the table and began to walk round her garden. There is
nothing like it for assisting thought.  However, to-day it was not
helping much; she went three times round and still couldn't think of a
rhyme for Hyacinth.  "Plinth" was a little difficult to work in;
"besides," she reminded herself, "I don't quite know what it means."
Belvane felt as I do about poetry: that however incomprehensible it
may be to the public, the author should be quite at ease with it.
</P>

<P>
She added up the lines she had written already&mdash;seventeen.  If she
stopped there, it would be the only epic that had stopped at the
seventeenth line.
</P>

<P>
She sighed, stretched her arms, and looked up at the sky. The weather
was all against her.  It was the ideal largesse morning. . . .
</P>

<P>
Twenty minutes later she was on her cream-white palfrey. Twenty-one
minutes later Henrietta Crossbuns had received a bag of gold neatly
under the eye, as she bobbed to her Ladyship.  To this extent only did
H. Crossbuns leave her mark upon Euralian history; but it was a mark
which lasted for a full month.
</P>

<P>
Hyacinth knew nothing of all this.  She did not even know that Belvane
was entering for the prize poem.  She had forgotten her promise to
encourage literature in the realm.
</P>

<P>
And why?  Ah, ladies, can you not guess why?  She was thinking of
Prince Udo of Araby.  What did he look like?  Was he dark or fair?
Did his hair curl naturally or not?
</P>

<P>
Was he wondering at all what <i>she</i> looked like?
</P>

<P>
Wiggs had already decided that he was to fall in love with her Royal
Highness and marry her.
</P>

<P>
"I think," said Wiggs, "that he'll be very tall, and have lovely blue
eyes and golden hair."
</P>

<P>
This is what they were like in all the books she had ever dusted; like
this were the seven Princes (now pursuing perilous adventures in
distant countries) to whom the King had promised Hyacinth's
hand&mdash;Prince Hanspatch of Tregong, Prince Ulric, the Duke of
Highanlow, and all the rest of them.  Poor Prince Ulric!  In the
moment of victory he was accidentally fallen upon by the giant whom he
was engaged in undermining, and lost all appetite for adventure
thereby.  Indeed, in his latter years he was alarmed by anything
larger than a goldfish, and lived a life of strictest seclusion.
</P>

<P class="noindent" align="center">
<a name="img0186"></a>
<img src="images/0186.jpg"
alt="[Illustration: Twenty-one minutes later Henrietta Crossbuns was acknowledging a bag of gold]">
<img src="images/0187.jpg"
alt="[Illustration: Twenty-one minutes later Henrietta Crossbuns was acknowledging a bag of gold]">
</P>

<P>
"<i>I</i> think he'll be dark," said Hyacinth.  Her own hair was
corn-coloured.
</P>

<P>
Poor Prince Hanspatch of Tregong; I've just remembered about him&mdash;no,
I haven't, it was the Duke of Highanlow. Poor Duke of Highanlow!  A
misunderstanding with a wizard having caused his head to face the
wrong way round, he was so often said good-bye to at the very moment
of arrival, that he gradually lost his enthusiasm for social
enterprises and confined himself to his own palace, where his
acrobatic dexterity in supplying himself with soup was a constant
source of admiration to his servants. . . .
</P>

<P>
However, it was Prince Udo of whom they were thinking now.  The
Messenger had returned from Araby; his Royal Highness must be expected
on the morrow.
</P>

<P>
"I do hope he'll be comfortable in the Purple Room," said Hyacinth.
"I wonder if it wouldn't have been better to have left him in the Blue
Room, after all."
</P>

<P>
They had had him in the Blue Room two days ago, until Hyacinth thought
that perhaps he would be more comfortable in the Purple Room, after
all.
</P>

<P>
"The Purple Room has the best view," said Wiggs helpfully.
</P>

<P>
"And it gets the sun.  Wiggs, don't forget to put some flowers there.
And have you given him any books?"
</P>

<P>
"I gave him two," said Wiggs.  "<i>Quests for Princes</i>, and <i>Wild
Animals at Home</i>."
</P>

<P>
"Oh, I'm sure he'll like those.  Now let's think what we shall do when
he comes.  He'll arrive some time in the afternoon.  Naturally he will
want a little refreshment."
</P>

<P>
"Would he like a picnic in the forest?" asked Wiggs.
</P>

<P>
"I don't think any one wants a picnic after a long journey."
</P>

<P>
"I <i>love</i> picnics."
</P>

<P>
"Yes, dear; but, you see, Prince Udo's much older than you, and I
expect he's had so many picnics that he's tired of them.  I suppose
really I ought to receive him in the Throne Room, but that's
so&mdash;so&mdash;&mdash;"
</P>

<P>
"Stuffy," said Wiggs.
</P>

<P>
"That's just it.  We should feel uncomfortable with each other the
whole time.  I think I shall receive him up here; I never feel so
nervous in the open air."
</P>

<P>
"Will the Countess be here?" asked Wiggs.
</P>

<P>
"No," said the Princess coldly.  "At least," she corrected herself,
"she will not be invited.  Good afternoon, Countess."  It was like
her, thought Hyacinth, to arrive at that very moment.
</P>

<P>
Belvane curtsied low.
</P>

<P>
"Good afternoon, your Royal Highness.  I am here purely on a matter of
business.  I thought it my duty to inform your Royal Highness of the
result of the Literature prize."  She spoke meekly, and as one who
forgave Hyacinth for her unkindness towards her.
</P>

<P>
"Certainly, Countess.  I shall be glad to hear."
</P>

<P>
The Countess unrolled a parchment.
</P>

<P>
"The prize has been won," she said, "by&mdash;&mdash;" she held the parchment a
little closer to her eyes, "by Charlotte Patacake."
</P>

<P>
"Oh, yes.  Who is she?"
</P>

<P>
"A most deserving woman, your Royal Highness.  If she is the woman I'm
thinking of, a most deserving person, to whom the money will be more
than welcome.  Her poem shows a sense of values combined
with&mdash;er&mdash;breadth, and&mdash;er&mdash;distance, such as I have seldom seen
equalled.  The&mdash;er&mdash;technique is only excelled by the&mdash;shall I
say?&mdash;tempermentality, the boldness of the colouring, by the&mdash;how
shall I put it?&mdash;the firmness of the outline.  In short&mdash;&mdash;"
</P>

<P>
"In short," said the Princess, "you like it."
</P>

<P>
"Your Royal Highness, it is unique.  But naturally you will wish to
hear it for yourself.  It is only some twelve hundred lines long.  I
will declaim it to your Royal Highness."
</P>

<P>
She held the manuscript out at the full length of her left arm, struck
an attitude with the right arm, and began in her most thrilling voice:
</P>

<P class="poem">
    "<i>King Merriwig the First rode out to war,</i><BR>
     &nbsp;<i>As many other kings&mdash;&mdash;</i>"<BR>
</P>

<P>
"Yes, Countess, but another time.  I am busy this afternoon.  As you
know, I think, the Prince Udo of Araby arrives to-morrow, and&mdash;&mdash;"
</P>

<P>
Belvane's lips were still moving, and her right arm swayed up and
down.  "<i>What gladsome cheers assailed the balmy air!</i>" she murmured
to herself, and her hand when up to heaven.  "<i>They come from north,
from south</i>" (she pointed in the directions mentioned), "<i>from
everywhere.  No wight that stood&mdash;&mdash;</i>"
</P>

<P>
"He will be received privately up here by myself in the first place,
and afterwards&mdash;&mdash;"
</P>

<P>
"<i>Could gaze upon the sight unmoved, I wot</i>," whispered Belvane, and
placed her hand upon her breast to show that anyhow it had been too
much for <i>her</i>.  "<i>Why do they march so&mdash;&mdash;</i>  I beg your Royal
Highness's pardon.  I was so carried away by this wonderful poem.  I
do beg of your Royal Highness to read it."
</P>

<P>
The Princess waved the manuscript aside.
</P>

<P>
"I am not unmindful of the claims of literature, Countess, and I shall
certainly read the poem another time.  Meanwhile I can, I hope, trust
you to see that the prize is awarded to the rightful winner.  What I
am telling you now is that the Prince Udo is arriving to-morrow."
</P>

<P>
Belvane looked innocently puzzled.
</P>

<P>
"Prince Udo&mdash;Udo&mdash;would that be Prince Udo of Carroway, your Royal
Highness?  A tall man with three legs?"
</P>

<P>
"Prince Udo of Araby," said Hyacinth severely.  "I think I have
already mentioned him to your ladyship.  He will make a stay of some
months."
</P>

<P>
"But how <i>delightful</i>, your Royal Highness, to see a man again!  We
were all getting so dull together!  We want a man to wake us up a
little, don't we, Wiggs?  I will go and give orders about his room at
once, your Royal Highness.  You will wish him to be in the Purple
Room, of course?"
</P>

<P>
That settled it.
</P>

<P>
"He will be in the Blue Room," said Hyacinth decidedly.
</P>

<P>
"Certainly, your Royal Highness.  Fancy, Wiggs, a man again!  I will
go and see about it now, if I may have your Royal Highness's leave to
withdraw?"
</P>

<P>
A little mystified by Belvane's manner, Hyacinth inclined her head,
and the Countess withdrew.
</P>

<P>
</P>

<P>
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap11"></A>
<img src="images/0197X.jpg" alt="[Illustration: Detail of Udo as an animal]">
</p>

<H3 ALIGN="center">
CHAPTER XI
</H3>

<H3 ALIGN="center">
WATERCRESS SEEMS TO GO WITH THE EARS
</H3>

<P>
Wiggs gave a parting pat to the tablecloth and stood looking at it
with her head on one side.
</P>

<P>
"Now, then," she said, "have we got everything?"
</P>

<P>
"What about sardines?" said Woggs in her common way.  (I don't know
what she's doing in this scene at all, but Roger Scurvilegs insists on
it.)
</P>

<P>
"I don't think a <i>Prince</i> would like <i>sardines</i>," said Wiggs.
</P>

<P>
"If <i>I'd</i> been on a long journey, I'd <i>love</i> sardines.  It <i>is</i> a very
long journey from Araby, isn't it?"
</P>

<P>
"Awful long.  Why, it's taken him nearly a week.  Perhaps," she added
hopefully, "he's had something on the way."
</P>

<P>
"Perhaps he took some sandwiches with him," said Woggs, thinking that
this would be a good thing to do.
</P>

<P>
"What do you think he'll be like, Woggs?"
</P>

<P>
Woggs though for a long time.
</P>

<P>
"Like the King," she said.  "Only different," she added, as an
afterthought.
</P>

<P>
Up came the Princess for the fifth time that afternoon, all
excitement.
</P>

<P>
"Well," she said, "is everything ready?"
</P>

<P>
"Yes, your Royal Highness.  Except Woggs and me didn't quite know
about sardines."
</P>

<P>
The Princess laughed happily.
</P>

<P>
"I think there will be enough there for him.  It all looks very nice."
</P>

<P>
She turned round and discovered behind her the last person she wanted
to see just then.
</P>

<P>
The-last-person-she-wanted-to-see-just-then curtsied effectively.
</P>

<P>
"Forgive me, your Royal Highness," she said profusely, "but I thought
I had left Charlotte Patacake's priceless manuscript up here.  No;
evidently I was mistaken, your Royal Highness.  I will withdraw, your
Royal Highness, as I know your Royal Highness would naturally wish to
receive his Royal Highness alone."
</P>

<P>
Listening to this speech one is impressed with Woggs' method of
calling everybody "Mum."
</P>

<P>
"Not at all, Countess," said Hyacinth coldly.  "We would prefer you to
stay and help us receive his Royal Highness.  He is a little late, I
think."
</P>

<P>
Belvane looked unspeakably distressed.
</P>

<P>
"Oh, I do <i>hope</i> that nothing has happened to him on the way," she
exclaimed.  "I've an uneasy feeling that something may have occurred."
</P>

<P class="noindent" align="center">
<a name="img0200"></a>
<img src="images/0200.jpg" alt="[Illustration: Princess Hyacinth gave a shriek and faltered slowly backwards, verso]">
<img src="images/0201.jpg" alt="[Illustration: Princess Hyacinth gave a shriek and faltered slowly backwards, recto]">
</P>

<P>
"What could have happened to him?" asked Hyacinth, not apparently very
much alarmed.
</P>

<P>
"Oh, your Royal Highness, it's just a sort of silly feeling of mine.
There may be nothing in it."
</P>

<P>
There was a noise of footsteps from below; a man's voice was heard.
The Princess and the Countess, both extremely nervous, but from
entirely different reasons, arranged suitable smiles of greeting upon
their faces; Wiggs and Woggs stood in attitudes of appropriate
meekness by the table.  The Court Painter could have made a beautiful
picture of it.
</P>

<P>
"His Royal Highness Prince Udo of Araby," announced the voice of an
attendant.
</P>

<P>
"A nervous moment," said Belvane to herself.  "Can the ring have
failed to act?"
</P>

<P>
Udo trotted in.
</P>

<P>
"It hasn't," said Belvane.
</P>

<P>
Princess Hyacinth gave a shriek, and faltered slowly backwards; Wiggs,
who was familiar with these little accidents in the books which she
dusted, and Woggs, who had a natural love for any kind of animal,
stood their ground.
</P>

<P>
"Whatever is it?" murmured Hyacinth.
</P>

<P>
It was as well that Belvane was there.
</P>

<P>
"Allow me to present to your Royal Highness," she said, stepping
forward, "his Royal Highness Prince Udo of Araby."
</P>

<P>
"Prince <i>Udo?</i>" said Hyacinth, all unwilling to believe it.
</P>

<P>
"I'm afraid so," said Udo gloomily.  He had thought over this meeting
a good deal in the last two or three days, and he realised now that he
had underestimated the difficulties of it.
</P>

<P>
Hyacinth remembered that she was a Princess and a woman.
</P>

<P>
"I'm delighted to welcome your Royal Highness to Euralia," she said.
"Won't you sit down&mdash;I mean up&mdash;er, down."  (How <i>did</i> rabbits sit?
Or whatever he was?)
</P>

<P>
Udo decided to sit up.
</P>

<P>
"Thank you.  You've no idea how difficult it is to talk on four legs
to somebody higher up.  It strains the neck so."
</P>

<P>
There was an awkward silence.  Nobody quite knew what to say.
</P>

<P>
Except Belvane.
</P>

<P>
She turned to Udo with her most charming smile.  "Did you have a
pleasant journey?" she asked sweetly.
</P>

<P>
"No," said Udo coldly.
</P>

<P>
"Oh, do tell us what happened to you?" cried Hyacinth. "Did you meet
some terrible enchanter on the way?  Oh, I am so dreadfully sorry."
</P>

<P>
When one is not feeling very well there is a certain type of question
which is always annoying.
</P>

<P>
"Can't you <i>see</i> what's happened to me?" said Udo crossly.  "I don't
know <i>how</i> it happened.  I had come two days' journey from Araby,
when&mdash;&mdash;"
</P>

<P>
"Please, your Royal Highness," said Wiggs, "is this <i>your</i> tail in the
salt?"  She took it out, gave it a shake, and handed it back to him.
</P>

<P>
"Oh, thank you, thank you&mdash;two days' journey from Araby when I woke up
one afternoon and found myself like this.  I ask you to imagine my
annoyance.  My first thought naturally was to return home and hide
myself; but I told myself, Princess, that <i>you</i> wanted me."
</P>

<P>
The Princess could not help being touched by this, said as it was with
a graceful movement of the ears and a caressing of the right whisker,
but she wondered a little what she would do with him now that she had
got him.
</P>

<P>
"Er&mdash;what <i>are</i> you?" put in Belvane kindly, knowing how men are
always glad to talk about themselves.
</P>

<P>
Udo had caught sight of a well-covered table, and was looking at it
with a curious mixture of hope and resignation.
</P>

<P>
"Very, very hungry," he said, speaking with the air of one who knows.
</P>

<P>
The Princess, whose mind had been travelling, woke up suddenly.
</P>

<P>
"Oh, I was forgetting my manners," she said with a smile for which the
greediest would have forgiven her.  "Let us sit down and refresh
ourselves.  May I present to your Royal Highness the Countess
Belvane."
</P>

<P>
"Do I shake hands or pat him?" murmured that mistress of Court
etiquette, for once at a loss.
</P>

<P>
Udo placed a paw over his heart and bowed profoundly.
</P>

<P>
"Charmed," he said gallantly, and coming from a cross between a lion,
a rabbit, and a woolly lamb the merest suggestion of gallantry has a
most pleasing effect.
</P>

<P>
They grouped themselves round the repast.
</P>

<P>
"A little sherbet, your Royal Highness?" said Hyacinth, who presided
over the bowl.
</P>

<P>
Udo was evidently longing to say yes, but hesitated.
</P>

<P>
"I wonder if I dare."
</P>

<P>
"It's very good sherbet," said Wiggs, to encourage him.
</P>

<P>
"I'm sure it is, my dear.  But the question is, Do I like sherbet?"
</P>

<P>
"You can't help knowing if you like <i>sherbet</i>."
</P>

<P>
"Don't bother him, Wiggs," said Hyacinth, "a venison sandwich, dear
Prince?"
</P>

<P>
"The question is, Do I like venison sandwiches?"
</P>

<P>
"<i>I</i> do," announced Woggs to any one who was interested.
</P>

<P>
"You see," explained Udo, "I really don't know <i>what</i> I like."
</P>

<P>
They were all surprised at this, particularly Woggs. Belvane, who was
enjoying herself too much to wish to do anything but listen, said
nothing, and it was the Princess who obliged Udo by asking him what he
meant.  It was a subject upon which he was longing to let himself go
to somebody.
</P>

<P>
"Well," he said, expanding himself a little, so that Wiggs had to
remove his tail this time from the custard, "what am I?"
</P>

<P>
Nobody ventured to offer an opinion.
</P>

<P>
"Am I a hare?  Then put me next to the red currant jelly, or whatever
it is that hares like."
</P>

<P>
The anxious eye of the hostess wandered over the table.
</P>

<P>
"Am I a lion?" went on Udo, developing his theme.  "Then pass me
Wiggs."
</P>

<P>
"Oh, please don't be a lion," said Wiggs gently, as she stroked his
mane.
</P>

<P>
"But haven't you a feeling for anything?" asked Hyacinth.
</P>

<P>
"I have a great feeling of emptiness.  I yearn for <i>something</i>, only I
don't quite know what."
</P>

<P>
"I hope it isn't sardines," whispered Wiggs to Woggs.
</P>

<P>
"But what have you been eating on the way?" asked the Princess.
</P>

<P>
"Oh, grass and things chiefly.  I thought I should be safe with
grass."
</P>

<P>
"And were you&mdash;er&mdash;safe?" asked Belvane, with a great show of anxiety.
</P>

<P>
Udo coughed and said nothing.
</P>

<P>
"I know it's silly of me," said Hyacinth, "but I still don't quite
understand.  I should have thought that if you were a&mdash;a&mdash;&mdash;"
</P>

<P>
"Quite so," said Udo.
</P>

<P>
"&mdash;then you would have known by instinct what a&mdash;a&mdash;&mdash;"
</P>

<P>
"Exactly," said Udo.
</P>

<P>
"Likes to eat."
</P>

<P>
"Ah, I thought you'd think that.  That's just what I thought when
this&mdash;when I began to feel unwell.  But I've worked it out since, and
it's all wrong."
</P>

<P>
"This <i>is</i> interesting," said Belvane, settling herself more
comfortably.  "<i>Do</i> go on."
</P>

<P>
"Well, when&mdash;&mdash;"  He coughed and looked round at them coyly.  "This is
really rather a delicate subject."
</P>

<P>
"Not at all," murmured Hyacinth.
</P>

<P>
"Well, it's like this.  When an enchanter wants to annoy you he
generally turns you into an animal of some kind."
</P>

<P>
Belvane achieved her first blush since she was seventeen.
</P>

<P>
"It <i>is</i> a humorous way they have," she said.
</P>

<P>
"But suppose you really were an animal altogether, it wouldn't annoy
you at all.  An elephant isn't annoyed at being an elephant; he just
tries to be a good elephant, and he'd be miserable if he couldn't do
things with his trunk. The annoying thing is to look like an elephant,
to have the very complicated&mdash;er&mdash;inside of an elephant, and yet all
the time really to be a man."
</P>

<P>
They were all intensely interested.  Woggs thought that it was going
to lead up to a revelation of what sort of animal Prince Udo really
was, but in this she was destined to be disappointed.  After all there
were advantages in Udo's present position.  As a man he had never been
listened to so attentively.
</P>

<P>
"Now suppose for a moment I am a lion.  I have the&mdash;er&mdash;delicate
apparatus of a lion, but the beautiful thoughts and aspirations of a
Prince.  Thus there is one&mdash;er&mdash;side of me which craves for raw beef,
but none the less there is a higher side of me" (he brought his paw up
towards his heart), "which&mdash;well, you know how <i>you'd</i> feel about it
yourself."
</P>

<P>
The Princess shuddered.
</P>

<P>
"I <i>should</i>," she said, with conviction.
</P>

<P>
Belvane was interested, but thought it all a little crude.
</P>

<P>
"You see the point," went on Udo.  "A baby left to itself doesn't know
what is good for it.  Left to itself it would eat anything.  Now turn
a man suddenly into an animal and he is in exactly the same state as
that baby."
</P>

<P>
"I hadn't thought of it like that," said Hyacinth.
</P>

<P>
"I've <i>had</i> to think of it!  Now let us proceed further with the
matter."  Udo was thoroughly enjoying himself.  He had not had such a
time since he had given an address on Beetles to all the leading
citizens of Araby at his coming-of-age.  "Suppose again that I am a
lion.  I know from what I have read or seen that raw meat agrees best
with the lion's&mdash;er&mdash;organisation, and however objectionable it might
look I should be foolish not to turn to it for sustenance.  But if you
don't quite know what animal you're supposed to be, see how difficult
the problem becomes.  It's a question of trying all sorts of horrible
things in order to find out what agrees with you."  His eyes took on a
faraway look, a look in which the most poignant memories seem to be
reflected.  "I've been experimenting," he said, "for the last three
days."
</P>

<P>
They all gazed sadly and sympathetically at him.  Except Belvane.  She
of course wouldn't.
</P>

<P>
"What went best?" she asked brightly.
</P>

<P>
"Oddly enough," said Udo, cheering up a little, "banana fritters.
Have you ever kept any animal who lived entirely on banana fritters?"
</P>

<P>
"Never," smiled the Princess.
</P>

<P>
"Well, that's the animal I probably am."  He sighed and added, "There
were one or two animals I wasn't."  For a little while he seemed to be
revolving bitter memories, and then went on, "I don't suppose any of
you here have any idea how very prickly thistles are when they are
going down.  Er&mdash;may I try a watercress sandwich?  It doesn't suit the
tail, but it seems to go with the ears."  He took a large bite and
added through the leaves, "I hope I don't bore you, Princess, with my
little troubles."
</P>

<P>
Hyacinth clasped his paw impulsively.
</P>

<P>
"My dear Prince Udo, I'm only longing to help.  We must think of some
way of getting this horrible enchantment off you.  There are so many
wise books in the library, and my father has composed a spell
which&mdash;oh, I'm sure we shall soon have you all right again."
</P>

<P>
Udo took another sandwich.
</P>

<P>
"Very good of you, Princess, to say so.  You understand how annoying a
little indisposition of this kind is to a man of my temperament."  He
beckoned to Wiggs.  "How do you make these?" he asked in an undertone.
</P>

<P>
Gracefully undulating, Belvane rose from her seat.
</P>

<P>
"Well," she said, "I must go and see that the stable&mdash;&mdash;" she broke
off in a pretty confusion&mdash;"How <i>silly</i> of me, I mean the Royal
Apartment is prepared.  Have I your Royal Highness's leave to
withdraw?"
</P>

<P>
She had.
</P>

<P>
"And, Wiggs, dear, you too had better run along and see if you can
help.  You may leave the watercress sandwiches," she added, as Wiggs
hesitated for a moment.
</P>

<P>
With a grateful look at her Royal Highness Udo helped himself to
another one.
</P>

<P>
</P>

<P>
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap12"></A>
<img src="images/0217X.jpg" alt="[Illustration: Detail of a child with a very large boot]">
</p>

<H3 ALIGN="center">
CHAPTER XII
</H3>

<H3 ALIGN="center">
WE DECIDE TO WRITE TO UDO'S FATHER
</H3>

<P>
"Now, my dear Princess," said Udo, as soon as they were alone.  "Let
me know in what way I can help you."
</P>

<P>
"Oh, Prince Udo," said Hyacinth earnestly, "it <i>is</i> so good of you to
have come.  I feel that this&mdash;this little accident is really my fault
for having asked you here."
</P>

<P>
"Not at all, dear lady.  It is the sort of little accident that might
have happened to anybody, anywhere.  If I can still be of assistance
to you, pray inform me.  Though my physical powers may not for the
moment be quite what they were, I flatter myself that my mental
capabilities are in no way diminished."  He took another bite of his
sandwich and wagged his head wisely at her.
</P>

<P>
"Let's come over here," said Hyacinth.
</P>

<P>
She moved across to an old stone seat in the wall, Udo following with
the plate, and made room for him by her side.  There is, of course, a
way of indicating to a gentleman that he may sit next to you on the
Chesterfield, and tell you what he has been doing in town lately, and
there is also another way of patting the sofa for Fido to jump up and
be-a-good-dog-and-lie-down-sir.  Hyacinth achieved something very
tactful in between, and Udo jumped up gracefully.
</P>

<P>
"Now we can talk," said Hyacinth.  "You noticed that lady, the
Countess Belvane, whom I presented to you?"
</P>

<P>
Udo nodded.
</P>

<P>
"What did you think of her?"
</P>

<P>
Udo was old enough to know what to say to that.
</P>

<P>
"I hardly looked at her," he said.  And he added with a deep bow,
"Naturally when your Royal Highness&mdash;oh, I beg your pardon, are my
ears in your way?"
</P>

<P>
"It's all right," said Hyacinth, rearranging her hair.  "Well, it was
because of that woman that I sent for you."
</P>

<P>
"But I can't marry her like this, your Royal Highness."
</P>

<P>
Hyacinth turned a startled face towards him.  Udo perceived that he
had blundered.  To hide his confusion he took another sandwich and ate
it very quickly.
</P>

<P>
"I want your help against her," said Hyacinth, a little distantly;
"she is plotting against me."
</P>

<P>
"Oh, your Royal Highness, now I see," said Udo, and he wagged his head
as much as to say, "You've come to the right man this time."
</P>

<P class="noindent" align="center">
<a name="img0220"></a>
<img src="images/0220.jpg"
alt="[Illustration: &quot;Now we can talk,&quot; said Hyacinth, verso]">
<img src="images/0221.jpg"
alt="[Illustration: &quot;Now we can talk,&quot; said Hyacinth, recto]">
</P>

<P>
"I don't trust her," said Hyacinth impressively.
</P>

<P>
"Well, now, Princess, I'm not surprised.  I'll tell you something
about that woman."
</P>

<P>
"Oh, what?"
</P>

<P>
"Well, when I was announced just now, what happened?  You, yourself,
Princess, were not unnaturally a little alarmed; those two little
girls were surprised and excited; but what of this Countess Belvane?
What did <i>she</i> do?"
</P>

<P>
"What <i>did</i> she do?"
</P>

<P>
"Nothing," said Udo impressively.  "She was neither surprised nor
alarmed."
</P>

<P>
"Why, now I come to think of it, I don't believe she was."
</P>

<P>
"And yet," said Udo half pathetically, half proudly, "Princes don't
generally look like this.  Now, why wasn't she surprised?"
</P>

<P>
Hyacinth looked bewildered.
</P>

<P>
"Did she know you were sending for me?" Udo went on.
</P>

<P>
"Yes."
</P>

<P>
"Because you had found out something about her?"
</P>

<P>
"Yes."
</P>

<P>
"Then depend upon it, <i>she's</i> done it.  <i>What</i> a mind that woman must
have!"
</P>

<P>
"But how could she do it?" exclaimed Hyacinth.  "Of course it's just
the sort of thing she <i>would</i> do if she could."
</P>

<P>
Udo didn't answer.  He was feeling rather annoyed with Belvane, and
had got off his seat and was trotting up and down so as not to show
his feelings before a lady.
</P>

<P>
"How <i>could</i> she do it?" implored Hyacinth.
</P>

<P>
"Oh, she's in with some enchanter or somebody," said Udo impatiently
as he trotted past.
</P>

<P>
Suddenly he had an idea.  He stopped in front of her.
</P>

<P>
"If only I were <i>sure</i> I was a lion."
</P>

<P>
He tried to roar, exclaimed hastily that it was only a practice one,
and roared again.  "No, I don't think I'm a lion after all," he
admitted sadly.
</P>

<P>
"Well," said Hyacinth, "we must think of a plan."
</P>

<P>
"We must think of a plan," said Udo, and he came and sat meekly beside
her again.  He could conceal it from himself no longer that he was not
a lion.  The fact depressed him.
</P>

<P>
"I suppose I have been weak," went on Hyacinth, "but ever since the
men went away she has been the ruling spirit of the country.  I think
she is plotting against me; I <i>know</i> she is robbing me.  I asked you
here so that you could help me to find her out."
</P>

<P>
Udo nodded his head importantly.
</P>

<P>
"We must watch her," he announced.
</P>

<P>
"We must watch her," agreed Hyacinth.  "It may take months&mdash;&mdash;"
</P>

<P>
"Did you say months?" said Udo, turning to her excitedly.
</P>

<P>
"Yes, why?"
</P>

<P>
"Well, it's&mdash;&mdash;" he gave a deprecating little cough.  "I know it's
very silly of me but&mdash;oh, well, let's hope it will be all right."
</P>

<P>
"Why, whatever is the matter?"
</P>

<P>
Udo was decidedly embarrassed.  He wriggled.  He drew little circles
with his hind paw on the ground and he shot little coy glances at her.
</P>

<P>
"Well, I"&mdash;and he gave a little nervous giggle&mdash;"I have a sort of
uneasy feeling that I may be one of those animals"&mdash;he gave another
conscious little laugh&mdash;"that have to go to sleep all through the
winter.  It would be very annoying&mdash;if I"&mdash;his paw became very busy
here&mdash;"if I had to dig a little hole in the ground, just when the plot
was thickening."
</P>

<P>
"Oh, but you won't," said Hyacinth, in distress.
</P>

<P>
They were both silent for a moment, thinking of the awful
possibilities.  Udo's tail had fallen across Hyacinth's lap, and she
began to play with it absently.
</P>

<P>
"Anyway," she said hopefully, "it's only July now."
</P>

<P>
"Ye&mdash;es," said Udo.  "I suppose I should get&mdash;er&mdash;busy about November.
We ought to find out something before then.  First of all we'd
better&mdash;&mdash;  Oh!"  He started up in dismay.  "I've just had a
<i>horrible</i> thought.  Don't I have to collect a little store of nuts
and things?"
</P>

<P>
"Surely&mdash;&mdash;"
</P>

<P>
"I should have to start that pretty soon," said Udo thoughtfully.
"You know, I shouldn't be very handy at it.  Climbing about after
nuts," he went on dreamily, "what a life for a&mdash;&mdash;"
</P>

<P>
"Oh, don't!" pleaded Hyacinth.  "Surely only squirrels do that?"
</P>

<P>
"Yes&mdash;yes.  Now, if I were a squirrel.  I should&mdash;may I have my tail
for a moment?"
</P>

<P>
"Oh, I'm so sorry," said Hyacinth, very much confused as she realised
the liberty she had been taking, and she handed his tail back to him.
</P>

<P>
"Not at all," said Udo.
</P>

<P>
He took it firmly in his right hand.  "Now then," he said, "we shall
see.  Watch this."
</P>

<P>
Sitting on his back legs he arched his tail over his head, and letting
go of it suddenly, began to nibble at a sandwich held in his two front
paws. . . .
</P>

<P>
A pretty picture for an artist.
</P>

<P>
But a bad model.  The tail fell with a thud to the ground.
</P>

<P>
"There!" said Udo triumphantly.  "That proves it.  I'm <i>not</i> a
squirrel."
</P>

<P>
"Oh, I'm so glad," said Hyacinth, completely convinced, as any one
would have been, by this demonstration.
</P>

<P>
"Yes, well, that's all right then.  Now we can make our plans.  First
of all we'd better&mdash;&mdash;"  He stopped suddenly, and Hyacinth saw that he
was gazing at his tail.
</P>

<P>
"Yes?" she said encouragingly.
</P>

<P>
He picked up his tail and held it out in front of him.  There was a
large knot in the middle of it.
</P>

<P>
"Now, <i>what</i> have I forgotten?" he said, rubbing his head
thoughtfully.
</P>

<P>
Poor Hyacinth!
</P>

<P>
"Oh, dear Prince Udo, I'm so sorry.  I'm afraid I did that without
thinking."
</P>

<P>
Udo, the gallant gentleman, was not found wanting.
</P>

<P>
"A lover's knot," he said, with a graceful incli&mdash;no, he stopped in
time.  But really, those ears of his made ordinary politeness quite
impossible.
</P>

<P>
"Oh, Udo," said Hyacinth impulsively, "if only I could help you to get
back to your proper form again."
</P>

<P>
"Yes, if only," said Udo, becoming practical again; "but how are we
going to do it?  Just one more watercress sandwich," he said
apologetically; "they go with the ears so well."
</P>

<P>
"I shall threaten the Countess," said Hyacinth excitedly.  "I shall
tell her that unless she makes the enchanter restore you to your
proper form, I shall put her in prison."
</P>

<P>
Udo was not listening.  He had gone off into his own thoughts.
"Banana fritters <i>and</i> watercress sandwiches," he was murmuring to
himself.  "I suppose I must be the only animal of the kind in the
world."
</P>

<P>
"Of course," went on Hyacinth, half to herself, "she might get the
people on her side, the ones that she's bribed. And if she did&mdash;&mdash;"
</P>

<P>
"That's all right, that's all right," said Udo grandly.  "Leave her to
me.  There's something about your watercress that inspires me to do
terrible deeds.  I feel a new&mdash;whatever I am."
</P>

<P>
One gathers reluctantly from this speech that Udo had partaken too
freely.
</P>

<P>
"Of course," said Hyacinth, "I could write to my father, who might
send some of his men back, but I shouldn't like to do that.  I
shouldn't like him to think that I had failed him."
</P>

<P>
"Extraordinary how I take to these things," said Udo, allowing himself
a little more room on the seat.  "Perhaps I am a rabbit after all.  I
wonder what I should look like behind wire netting."  He took another
bite and went on, "I wonder what I should do if I saw a ferret.  I
suppose you haven't got a ferret on you, Princess?"
</P>

<P>
"I beg your pardon, Prince?  I'm afraid I was thinking of something
else.  What did you say?"
</P>

<P>
"Nothing, nothing.  One's thoughts run on."  He put his hand out for
the plate, and discovered that it was empty. He settled himself more
comfortably, and seemed to be about to sink into slumber when his
attention was attracted suddenly by the knot in his tail.  He picked
it up and began lazily to undo it.  "I wish I could lash my tail," he
murmured; "mine seems to be one of the tails that don't lash."  He
began very gingerly to feel the tip of it.  "I wonder if I've got a
sting anywhere."  He closed his eyes, muttering, "Sting Countess neck,
sting all over neck, sting lots stings," and fell peacefully asleep.
</P>

<P>
It was a disgraceful exhibition.  Roger Scurvilegs tries to slur it
over; talks about the great heat of the sun, and the notorious effect
of even one or two watercress sandwiches on an empty&mdash;on a man who has
had nothing to eat for several days.  This is to palter with the
facts.  The effect of watercress sandwiches upon Udo's arrangements
(however furnished) we have all just seen for ourselves; but what
Roger neglects to lay stress upon is the fact that it was the effect
of twenty-one or twenty-two watercress sandwiches.  There is no
denying that it was a disgraceful exhibition.  If I had been there, I
should certainly have written to his father about it.
</P>

<P>
Hyacinth looked at him uneasily.  Her first feeling was one of
sympathy.  "Poor fellow," she thought, "he's had a hard time lately."
But it is a strain on the sympathy to gaze too long on a mixture of
lion, rabbit, and woolly lamb, particularly when the rabbit part has
its mouth open and is snoring gently.
</P>

<P>
Besides, what could she do with him?  She had two of them on her hands
now: the Countess and the Prince.  Belvane was in an even better
position than before.  She could now employ Udo to help her in her
plots against the Princess.  "Grant to me so and so, or I'll keep the
enchantment for ever on his Royal Highness."  And what could a poor
girl do?
</P>

<P>
Well, she would have to come to some decision in the future.
Meanwhile the difficulties of the moment were enough.  The most
obvious difficulty was his bedroom.  Was it quite the sort of room he
wanted now?  Hyacinth realised suddenly that to be hostess to such a
collection of animals as Udo was would require all the tact she
possessed.  Perhaps he would tell her what he wanted when he woke up.
Better let him sleep peacefully now.
</P>

<P>
She looked at him, smiled in spite of herself, and went quickly down
into the Palace.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap13"></A>
<img src="images/0235X.jpg" alt="[Illustration: Detail of Belvane with castle in the background]">
</p>

<H3 ALIGN="center">
CHAPTER XIII
</H3>

<H3 ALIGN="center">
"PINK" RHYMES WITH "THINK"
</H3>

<P>
Udo awoke, slightly refreshed, and decided to take a firm line with
the Countess at once.  He had no difficulty about finding his way down
to her.  The Palace seemed to be full of servants, all apparently busy
about something which brought them for a moment in sight of the newly
arrived Prince, and then whisked them off, hand to mouth and shoulders
shaking.  By one of these, with more control over her countenance than
the others, an annoyed Udo was led into Belvane's garden.
</P>

<P>
She was walking up and down the flagged walk between her lavender
hedges, and as he came in she stopped and rested her elbows on her
sundial, and looked mockingly at him, waiting for him to speak.
"Between the showers I mark the hours," said the sundial (on the
suggestion of Belvane one wet afternoon), but for the moment the
Countess was in the way.
</P>

<P>
"Ah, here we are," said Udo in rather a nasty voice.
</P>

<P>
"Here we are," said Belvane sweetly.  "All of us."
</P>

<P>
Suddenly she began to laugh.
</P>

<P>
"Oh, Prince Udo," she said, "you'll be the death of me.  Count me as
one more of your victims."
</P>

<P>
It is easy to be angry with any one who will laugh at you all the
time, but difficult to be effective; particularly when&mdash;but we need
not dwell upon Udo's handicap again.
</P>

<P>
"I don't see anything to laugh at," he said stiffly.  "To intelligent
people the outside appearance is not everything."
</P>

<P>
"But it can be very funny, can't it?" said Belvane coaxingly.  "I
wished for something humorous to happen to you, but I never
thought&mdash;&mdash;"
</P>

<P>
"Ah," said Udo, "now we've got it."
</P>

<P>
He spoke with an air of a clever cross-examiner who has skilfully
extracted an admission from a reluctant witness.  This sort of tone
goes best with one of those keen legal faces; perhaps that is why
Belvane laughed again.
</P>

<P>
"You practically confess that you did it," went on Udo magnificently.
</P>

<P>
"Did what?"
</P>

<P>
"Turned me into a&mdash;a&mdash;&mdash;"
</P>

<P>
"A rabbit?" said Belvane innocently.
</P>

<P>
A foolish observation like this always pained Udo.
</P>

<P>
"What makes you think I'm a rabbit?" he asked.
</P>

<P>
"I don't mind what you are, but you'll never dare show yourself in the
country like this."
</P>

<P>
"Be careful, woman; don't drive me too far.  Beware lest you rouse the
lion in me."
</P>

<P>
"Where?" asked Belvane, with a child-like air.
</P>

<P>
With a gesture full of dignity and good breeding Udo called attention
to his tail.
</P>

<P>
"That," said the Countess, "is not the part of the lion that I'm
afraid of."
</P>

<P>
For the moment Udo was nonplussed, but he soon recovered himself.
</P>

<P>
"Even supposing&mdash;just for the sake of argument&mdash;that I am a rabbit, I
still have something up my sleeve; I'll come and eat your young
carnations."
</P>

<P>
Belvane adored her garden, but she was sustained by the thought that
it was only July just now.  She pointed this out to him.
</P>

<P>
"It needn't necessarily be carnations," he warned her.
</P>

<P>
"I don't want to put my opinion against one who has (forgive me)
inside knowledge on the subject, but I think I have nothing in my
garden at this moment that would agree with a rabbit."
</P>

<P>
"I don't mind if it <i>doesn't</i> agree with me," said Udo heroically.
</P>

<P>
This was more serious.  Her dear garden in which she composed, ruined
by the mastications&mdash;machinations&mdash;what was the word?&mdash;of an enemy!
The thought was unbearable.
</P>

<P>
"You aren't a rabbit," she said hastily; "you aren't really a rabbit.
Because&mdash;because you don't <i>woffle</i> your nose properly."
</P>

<P>
"I could," said Udo simply.  "I'm just keeping it back, that's all."
</P>

<P>
"Show me how," cried Belvane, clasping her hands eagerly together.
</P>

<P>
It was not what he had come into the garden for, and it accorded ill
with the dignity of the Royal House of Araby, but somehow one got led
on by this wicked woman.
</P>

<P>
"Like this," said Udo.
</P>

<P>
The Countess looked at him critically with her head on one side.
</P>

<P>
"No," she said, "that's quite wrong."
</P>

<P>
"Naturally I'm a little out of practice."
</P>

<P>
"I'm sorry," said Belvane.  "I'm afraid I can't pass you."
</P>

<P>
Udo couldn't think what had happened to the conversation.  With a
great effort he extracted himself from it.
</P>

<P>
"Enough of this, Countess," he said sternly.  "I have your admission
that it was you who put this enchantment on me."
</P>

<P>
"It was I.  I wasn't going to have you here interfering with my
plans."
</P>

<P>
"Your plans to rob the Princess."
</P>

<P>
Belvane felt that it was useless to explain the principles of
largesse-throwing to Udo.  There will always be men like Udo and Roger
Scurvilegs who take these narrow matter-of-fact views.  One merely
wastes time in arguing with them.
</P>

<P>
"My plans," she repeated.
</P>

<P>
"Very well.  I shall go straight to the Princess, and she will unmask
you before the people."
</P>

<P>
Belvane smiled happily.  One does not often get such a chance.
</P>

<P>
"And who," she asked sweetly, "will unmask your Royal Highness before
the people, so that they may see the true Prince Udo underneath?"
</P>

<P>
"What do you mean?" said Udo, though he was beginning to guess.
</P>

<P>
"That noble handsome countenance which is so justly the pride of
Araby&mdash;how shall we show that to the people? They'll form such a
mistaken idea of it if they all see you like this, won't they?"
</P>

<P>
Udo was quite sure now that he understood.  Hyacinth had understood at
the very beginning.
</P>

<P class="noindent" align="center">
<a name="img0242"></a>
<img src="images/0242.jpg" alt="[Illustration: He forgot his manners, and made a jump towards her]">
<a name="img0243"></a>
<img src="images/0243.jpg" alt="[Illustration: She glided gracefully behind the sundial in a pretty affectation of alarm]">
</P>

<P>
"You mean that if the Princess Hyacinth falls in with your plans, you
will restore me to my proper form, but that otherwise you will leave
me like this?"
</P>

<P>
"One's actions are very much misunderstood," sighed Belvane.  "I've no
doubt that that is how it will appear to future historians."
</P>

<P>
(To Roger, certainly.)
</P>

<P>
It was too much for Udo.  He forgot his manners and made a jump
towards her.  She glided gracefully behind the sundial in a pretty
affectation of alarm . . . and the next moment Udo decided that the
contest between them was not to be settled by such rough-and-tumble
methods as these.  The fact that his tail had caught in something
helped him to decide.
</P>

<P>
Belvane was up to him in an instant.
</P>

<P>
"There, there!" she said soothingly, "Let <i>me</i> undo it for your Royal
Highness."  She talked pleasantly as she worked at it.  "Every little
accident teaches us something.  Now if you'd been a rabbit this
wouldn't have happened."
</P>

<P>
"No, I'm not even a rabbit," said Udo sadly.  "I'm just nothing."
</P>

<P>
Belvane stood up and made him a deep curtsey.
</P>

<P>
"You are his Royal Highness Prince Udo of Araby.  Your Royal
Highness's straw is prepared.  When will your Royal Highness be
pleased to retire?"
</P>

<P>
It was a little unkind, I think.  I should not record it of her were
not Roger so insistent.
</P>

<P>
"Now," said Udo, and lolloped sadly off.  It was his one really
dignified moment in Euralia.
</P>

<P>
On his way to his apartment he met Wiggs.
</P>

<P>
"Wiggs," he said solemnly, "if ever you can do anything to annoy that
woman, such as making her an apple-pie bed, or <i>anything</i> like that, I
wish you'd do it."
</P>

<P>
Whereupon he retired for the night.  Into the mysteries of his toilet
we had perhaps better not inquire.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
As the chronicler of these simple happenings many years ago, it is my
duty to be impartial.  "These are the facts," I should say, "and it is
for your nobilities to judge of them.  Thus and thus my characters
have acted; how say you, my lords and ladies?"
</P>

<P>
I confess that this attitude is beyond me; I have a fondness for all
my people, and I would not have you misunderstand any of them.  But
with regard to one of them there is no need for me to say anything in
her defence.  About her at any rate we agree.
</P>

<P>
I mean Wiggs.  We take the same view as Hyacinth: she was the best
little girl in Euralia.  It will come then as a shock to you (as it
did to me on the morning after I had staggered home with Roger's
seventeen volumes) to learn that on her day Wiggs could be as bad as
anybody.  I mean really bad.  To tear your frock, to read books which
you ought to be dusting, these are accidents which may happen to
anybody.  Far otherwise was Wiggs's fall.
</P>

<P>
She adopted, in fact, the infamous suggestion of Prince Udo.  Three
nights later, with malice aforethought and to the comfort of the
King's enemies and the prejudice of the safety of the realm, she made
an apple-pie bed for the Countess.
</P>

<P>
It was the most perfect apple-pie bed ever made.  Cox himself could
not have improved upon it; Newton has seen nothing like it.  It took
Wiggs a whole morning; and the results, though private (that is the
worst of an apple-pie bed), were beyond expectation.  After wrestling
for half an hour the Countess spent the night in a garden hammock,
composing a bitter Ode to Melancholy.
</P>

<P>
Of course Wiggs caught it in the morning; the Countess suspected what
she could not prove.  Wiggs, now in for a thoroughly bad week,
realised that it was her turn again.  What should she do?
</P>

<P>
An inspiration came to her.  She had been really bad the day before;
it was a pity to waste such perfect badness as that.  Why not have the
one bad wish to which the ring entitled her?
</P>

<P>
She drew the ring out from its hiding-place round her neck.
</P>

<P>
"I wish," she said, holding it up, "I wish that the Countess
Belvane&mdash;&mdash;" she stopped to think of something that would really annoy
her&mdash;"I wish that the Countess shall never be able to write another
rhyme again."
</P>

<P>
She held her breath, expecting a thunderclap or some other outward
token of the sudden death of Belvane's muse. Instead she was struck by
the extraordinary silence of the place.  She had a horrid feeling that
everybody else was dead, and realising all at once that she was a very
wicked little girl, she ran up to her room and gave herself up to
tears.
</P>

<P>
MAY YOU, DEAR SIR OR MADAM, REPENT AS QUICKLY!
</P>

<P>
However, this is not a moral work.  An hour later Wiggs came into
Belvane's garden, eager to discover in what way her inability to rhyme
would manifest itself.  It seemed that she had chosen the exact
moment.
</P>

<P>
In the throes of composition Belvane had quite forgotten the apple-pie
bed, so absorbing is our profession.  She welcomed Wiggs eagerly, and
taking her hand led her towards the roses.
</P>

<P>
"I have just been talking to my dear roses," she said.  "Listen:
</P>

<P class="poem">
    <i>Whene'er I take my walks about,</i> <BR>
    <i>I like to see the roses out;</i><BR>
    <i>I like them yellow, white, and pink,</i> <BR>
    <i>But crimson are the best, I think.</i> <BR>
    <i>The butterfly&mdash;&mdash;</i>"<BR>
</P>

<P>
But we shall never know about the butterfly.  It may be that Wiggs has
lost us here a thought on lepidoptera which the world can ill spare;
for she interrupted breathlessly.
</P>

<P>
"When did you write that?"
</P>

<P>
"I was just making it up when you came in, dear child.  These thoughts
often come to me as I walk up and down my beautiful garden.  '<i>The
butterfly&mdash;&mdash;</i>'"
</P>

<P>
But Wiggs had let go her hand and was running back to the Palace.  She
wanted to be alone to think this out.
</P>

<P>
What had happened?  That it was truly a magic ring, as the fairy had
told her, she had no doubt; that her wish was a bad one, that she had
been bad enough to earn it, she was equally certain.  What then had
happened?  There was only one answer to her question.  The bad wish
had been granted to someone else.
</P>

<P>
To whom?  She had lent the ring to nobody.  True, she had told the
Princess all about it, but&mdash;&mdash;
</P>

<P>
Suddenly she remembered.  The Countess had had it in her hands for a
moment.  Yes, and she had sent her out of the room, and&mdash;
</P>

<P>
So many thoughts crowded into Wiggs's mind at this moment that she
felt she must share them with somebody.  She ran off to find the
Princess.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<img src="images/0253X.jpg" alt="[Illustration: Detail of Wiggs curtsying]">
<A NAME="chap14"></A>
</p>

<H3 ALIGN="center">
CHAPTER XIV
</H3>

<H3 ALIGN="center">
"WHY CAN'T YOU BE LIKE WIGGS?"
</H3>

<P>
Hyacinth was with Udo in the library.  Udo spent much of his time in
the library nowadays; for surely in one of those many books was to be
found some Advice to a Gentleman in Temporary Difficulties suitable to
a case like his.  Hyacinth kept him company sadly.  It had been such a
brilliant idea inviting him to Euralia; how she wished now that she
had never done it.
</P>

<P>
"Well, Wiggs," she said, with a gentle smile, "what have you been
doing with yourself all the morning?"
</P>

<P>
Udo looked up from his mat and nodded to her.
</P>

<P>
"I've found out," said Wiggs excitedly; "it was the <i>Countess</i> who did
it."
</P>

<P>
Udo surveyed her with amazement.
</P>

<P>
"The Princess Hyacinth," he said, "has golden hair.  One discovers
these things gradually."  And he returned to his book.
</P>

<P>
Wiggs looked bewildered.
</P>

<P>
"He means, dear," said Hyacinth, "that it is quite obvious that the
Countess did it, and we have known about it for days."
</P>

<P>
Udo wore, as far as his face would permit, the slightly puffy
expression of one who has just said something profoundly ironical and
is feeling self-conscious about it.
</P>

<P>
"Oh&mdash;h," said Wiggs in such a disappointed voice that it seemed as if
she were going to cry.
</P>

<P>
Hyacinth, like the dear that she was, made haste to comfort her.
</P>

<P>
"We didn't really <i>know</i>," she said; "we only guessed it.  But now
that you have found out, I shall be able to punish her properly.  No,
don't come with me," she said, as she rose and moved towards the door;
"stay here and help his Royal Highness.  Perhaps you can find the book
that he wants; you've read more of them than I have, I expect."
</P>

<P>
Left alone with the Prince, Wiggs was silent for a little, looking at
him rather anxiously.
</P>

<P>
"Do you know <i>all</i> about the Countess?" she asked at last.
</P>

<P>
"If there's anything I don't know, it must be <i>very</i> bad."
</P>

<P>
"Then you know that it's all my fault that you are like this?  Oh,
dear Prince Udo, I am so dreadfully sorry."
</P>

<P>
"What do you mean&mdash;<i>your</i> fault?"
</P>

<P>
"Because it was my ring that did it."
</P>

<P>
Udo scratched his head in a slightly puzzled but quite a nice way.
</P>

<P>
"Tell me all about it from the beginning," he said.  "You have found
out something after all, I believe."
</P>

<P>
So Wiggs told her story from the beginning.  How the fairy had given
her a ring; how the Countess had taken it from her for five minutes
and had a bad wish on it; and how Wiggs had found her out that very
morning.
</P>

<P>
Udo was intensely excited by the story.  He trotted up and down the
library, muttering to himself.  He stopped in front of Wiggs as soon
as she had finished.
</P>

<P>
"Is the ring still going?" he asked.  "I mean, can you have another
wish on it?"
</P>

<P>
"Yes, just one."
</P>

<P>
"Then wish her to be turned into a&mdash;&mdash;"  He tried to think of
something that would meet the case.  "What about a spider?" he said
thoughtfully.
</P>

<P>
"But that's a <i>bad</i> wish," said Wiggs.
</P>

<P>
"Yes, but it's <i>her</i> turn."
</P>

<P>
"Oh, but I'm only allowed a good wish now."  She added rapturously,
"And I know what it's going to be."
</P>

<P>
So did Udo.  At least he thought he did.
</P>

<P>
"Oh, you dear," he said, casting an affectionate look on her.
</P>

<P>
"Yes, that's it.  That I might be able to dance like a fairy."
</P>

<P>
Udo could hardly believe his ears, and they were adequate enough for
most emergencies.
</P>

<P>
"But how is that going to help <i>me?</i>" he said, tapping his chest with
his paw.
</P>

<P>
"But it's <i>my</i> ring," said Wiggs.  "And so of course I'm going to wish
that I can dance like a fairy.  I've always meant to, as soon as I've
been good for a day first."
</P>

<P>
The child was absurdly selfish.  Udo saw that he would have to appeal
to her in another way.
</P>

<P>
"Of course," he began, "I've nothing to say against dancing <i>as</i>
dancing, but I think you'll get tired of it.  Just as I shall get
tired of&mdash;lettuce."
</P>

<P>
Wiggs understood now.
</P>

<P>
"You mean that I might wish you to be a Prince again?"
</P>

<P>
"Well," said Udo casually, "it just occurred to me as an example of
what might be called the Good Wish."
</P>

<P>
"Then I shall never be able to dance like a fairy?"
</P>

<P>
"Neither shall I, if it comes to that," said Udo.  Really, the child
was very stupid.
</P>

<P>
"Oh, it's too cruel," said Wiggs, stamping her foot.  "I did so want
to be able to dance."
</P>

<P>
Udo glanced gloomily into the future.
</P>

<P>
"To live for ever behind wire netting," he mused; "to be eternally
frightened by pink-eyed ferrets; to be offered
bran-mash&mdash;bran-mash&mdash;bran-mash wherever one visited week after week,
month after month, year after year, century after&mdash;how long <i>do</i>
rabbits live?"
</P>

<P>
But Wiggs was not to be moved.
</P>

<P>
"I <i>won't</i> give up my wish," she said passionately.
</P>

<P>
Udo got on to his four legs with dignity.
</P>

<P>
"Keep your wish," he said.  "There are plenty of other ways of getting
out of enchantments.  I'll learn up a piece of poetry by our Court
Poet Sacharino, and recite it backwards when the moon is new.
Something like that.  I can do this quite easily by myself.  Keep your
wish."
</P>

<P>
He went slowly out.  His tail (looking more like a bell-rope than
ever) followed him solemnly.  The fluffy part that you pull was for a
moment left behind; then with a jerk it was gone, and Wiggs was left
alone.
</P>

<P>
"I won't give up my wish," cried Wiggs again.  "I'll wish it now
before I'm sorry."  She held the ring up.  "I wish that&mdash;&mdash;"  She
stopped suddenly.  "Poor Prince Udo he seems very unhappy.  I wonder
if it <i>is</i> a good wish to wish to dance when people are unhappy."  She
thought this out for a little, and then made her great resolve.
"Yes," she said, "I'll wish him well again."
</P>

<P>
Once more she held the ring up in her two hands.
</P>

<P>
"I wish," she said, "that Prince Udo&mdash;&mdash;"
</P>

<P>
I know what you're going to say.  It was no good her wishing her good
wish, because she had been a bad girl the day before&mdash;making the
Countess an apple-pie bed and all&mdash;disgraceful!  How could she
possibly suppose&mdash;&mdash;
</P>

<P>
She didn't.  She remembered just in time.
</P>

<P>
"Oh, bother," said Wiggs, standing in the middle of the room with the
ring held above her head.  "I've got to be good for a day first.
<i>Bother!</i>"
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
So the next day was Wiggs's Good Day.  The legend of it was handed
down for years afterwards in Euralia.  It got into all the
Calendars&mdash;July 20th it was&mdash;marked with a red star; in Roger's
portentous volumes it had a chapter devoted to it.  There was some
talk about it being made into a public holiday, he tells us, but this
fell through.  Euralian mothers used to scold their naughty children
with the words, "Why can't you be like Wiggs?" and the children used
to tell each other that there never was a real Wiggs, and that it was
only a made-up story for parents.  However, you have my word for it
that it was true.
</P>

<P>
She began by getting up at five o'clock in the morning, and after
dressing herself very neatly (and being particularly careful to wring
out her sponge) she made her own bed and tidied up the room.  For a
moment she thought of waking the grown-ups in the Palace and letting
them enjoy the beautiful morning too, but a little reflection showed
her that this would not be at all a kindly act; so, having dusted the
Throne Room and performed a few simple physical exercises, she went
outside and attended to the smaller domestic animals.
</P>

<P class="noindent" align="center">
<a name="img0262"></a><img src="images/0262.jpg" alt="[Illustration: When anybody of superior station or age came into the
room she rose and curtsied, verso]">
<img src="images/0263.jpg" alt="[Illustration: When anybody of superior station or age came into the
room she rose and curtsied, recto]">
</P>

<P>
At breakfast she had three helps of something very nutritious, which
the Countess said would make her grow, but only one help of everything
else.  She sat up nicely all the time, and never pointed to anything
or drank with her mouth full.  After breakfast she scattered some
crumbs on the lawn for the robins, and then got to work again.
</P>

<P>
First she dusted and dusted and dusted; then she swept and swept and
swept; then she sewed and sewed and sewed.  When anybody of superior
station or age came into the room she rose and curtsied and stood with
her hands behind her back, while she was being spoken to.  When
anybody said, "I wonder where I put my so-and-so," she jumped up and
said, "Let <i>me</i> fetch it," even if it was upstairs.
</P>

<P>
After dinner she made up a basket of provisions and took them to the
old women who lived near the castle; to some of them she sang or read
aloud, and when at one cottage she was asked, "Now won't you give me a
little dance," she smiled bravely and said, "I'm afraid I don't dance
very well."  I think that was rather sweet of her; if I had been the
fairy I should have let her off the rest of the day.
</P>

<P>
When she got back to the Palace she drank two glasses of warm milk,
with the skin on, and then went and weeded the Countess's lawn; and
once when she trod by accident on a bed of flowers, she left the
footprint there instead of scraping it over hastily, and pretending
that she hadn't been near the place, as you would have done.
</P>

<P>
And at half-past six she kissed everybody good-night (including Udo)
and went to bed.
</P>

<P>
So ended July the Twentieth, perhaps the most memorable day in
Euralian history.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
Udo and Hyacinth spent the great day peacefully in the library.  A
gentleman for all his fur, Udo had not told the Princess about Wiggs's
refusal to help him.  Besides, a man has his dignity.  To be turned
into a mixture of three animals by a woman of thirty, and to be turned
back again by a girl of ten, is to be too much the plaything of the
sex.  It was time he did something for himself.
</P>

<P>
"Now then, how did that bit of Sacharino's go?  Let me see."  He beat
time with a paw.  "'Blood for something, something, some&mdash;&mdash;'
Something like that.  'Blood for&mdash;er&mdash;blood for&mdash;er&mdash;&mdash;'  No, it's
gone again.  I know there was a bit of blood in it."
</P>

<P>
"I'm sure you'll get it soon," said Hyacinth.  "It sounds as thought
it's going to be just the sort of thing that's wanted."
</P>

<P>
"Oh, I shall get it all right.  Some of the words have escaped me for
the moment, that's all.  'Blood&mdash;er&mdash;blood.'  You must have heard of
it, Princess: it's about blood for he who something; you must know the
one I mean.
</P>

<P>
"I know I've heard of it," said the Princess, wrinkling her forehead,
"only I can't quite think of it for the moment.  It's about a&mdash;a&mdash;&mdash;"
</P>

<P>
"Yes, that's it," said Udo.
</P>

<P>
Then they both looked up at the ceiling with their heads on one side
and murmured to themselves.
</P>

<P>
But noon came and still they hadn't thought of it.
</P>

<P>
After a simple meal they returned to the library.
</P>

<P>
"I think I'd better write to Coronel," said Udo, "and ask him about
it."
</P>

<P>
"I thought you said his name was Sacharino."
</P>

<P>
"Oh, this is not the poet, it's just a friend of mine, but he's rather
good at this sort of thing.  The trouble is that it takes such a long
time for a letter to get there and back."
</P>

<P>
At the word "letter," Hyacinth started suddenly.
</P>

<P>
"Oh, Prince Udo," she cried, "I can never forgive myself. I've just
remembered the very thing.  Father told me in his letter that a little
couplet he once wrote was being very useful for&mdash;er&mdash;removing things."
</P>

<P>
"What sort of things?" said Udo, not too hopefully.
</P>

<P>
"Oh, enchantments and things."
</P>

<P>
Udo was a little annoyed at the "and things"&mdash;as those turning him
back into a Prince again was as much in the day's work as removing
rust from a helmet.
</P>

<P>
"It goes like this," said Hyacinth.
</P>

<P class="poem">
    "<i>Bo, boll, bill, bole.</i> <BR>
     &nbsp;<i>Wo, woll, will, wole.</i>"<BR>
</P>

<P>
"It sounds as though it would remove <i>anything</i>," she added, with a
smile.
</P>

<P>
Udo sat up rather eagerly.
</P>

<P>
"I'll try," he said.  "Is there any particular action that goes with
it?"
</P>

<P>
"I've never heard of any.  I expect you ought to say it as if you
meant it."
</P>

<P>
Udo sat up on his back paws, and, gesticulating freely with his right
paw, declaimed:
</P>

<P class="poem">
    "<i>Bo, boll, bill, bole.</i> <BR>
     &nbsp;<i>Wo, woll, will, wole.</i>"<BR>
</P>

<P>
He fixed his eyes on his paws, waiting for the transformation.
</P>

<P>
He waited.
</P>

<P>
And waited.
</P>

<P>
Nothing happened.
</P>

<P>
"It must be all right," said Hyacinth anxiously, "because I'm sure
Father would know.  Try saying it more like this."
</P>

<P>
She repeated the lines in a voice so melting, yet withal so dignified,
that the very chairs might have been expected to get up and walk out.
</P>

<P>
Udo imitated her as well as he could.
</P>

<P>
At about the time when Wiggs was just falling asleep, he repeated it
in his fiftieth different voice.
</P>

<P>
"I'm sorry," said Hyacinth; "perhaps it isn't so good as Father
thought it was."
</P>

<P>
"There's just one chance," said Udo.  "It's possible it may have to be
said on an empty stomach.  I'll try it to-morrow before breakfast."
</P>

<P>
Upstairs Wiggs was dreaming of the dancing that she had given up for
ever.
</P>

<P>
And what Belvane was doing I really don't know.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap15"></A>
<img src="images/0271X.jpg" alt="[Illustration: Detail of Wiggs dancing]">
</p>

<H3 ALIGN="center">
CHAPTER XV
</H3>

<H3 ALIGN="center">
THERE IS A LOVER WAITING FOR HYACINTH
</H3>

<P>
So the next morning before breakfast Wiggs went up on to the castle
walls and wished.  She looked over the meadows, and across the
peaceful stream that wandered through them, to the forest where she
had met her fairy, and she gave a little sigh.  "Good-bye, dancing,"
she said; and then she held the ring up and went on bravely, "Please I
was a very good girl all yesterday, and I wish that Prince Udo may be
well again."
</P>

<P>
For a full minute there was silence.  Then from the direction of Udo's
room below there came these remarkable words:
</P>

<P>
"<i>Take the beastly stuff away, and bring me a beefsteak and a flagon
of sack!</i>"
</P>

<P>
Between smiles and tears Wiggs murmured, "He <i>sounds</i> all right.  I
<i>am</i> g&mdash;glad."
</P>

<P>
And then she could bear it no longer.  She hurried down and out of the
Palace&mdash;away, away from Udo and the Princess and the Countess and all
their talk, to the cool friendly forest, there to be alone and to
think over all that she had lost.
</P>

<P>
It was very quiet in the forest.  At the foot of her own favourite
tree, a veteran of many hundred summers who stood sentinel over an
open glade that dipped to a gurgling brook and climbed gently away
from it, she sat down.  On the soft green yonder she might have
danced, an enchanted place, and now&mdash;never, never, never. . . .
</P>

<P>
How long had she sat there?  It must have been a long time&mdash;because
the forest had been so quiet, and now it was so full of sound.  The
trees were murmuring something to her, and the birds were singing it,
and the brook was trying to tell it too, but it would keep chuckling
over the very idea so that you could hardly hear what it was saying,
and there were rustlings in the grass&mdash;"Get up, get up," everything
was calling to her; "dance, dance."
</P>

<P>
She got up, a little frightened.  Everything seemed so strangely
beautiful.  She had never felt it like this before.  Yes, she would
dance.  She must say, "Thank you," for all this somehow; perhaps they
would excuse her if it was not very well expressed.
</P>

<P>
"This will just be for 'Thank you'" she said as she got up.  "I shall
never dance again."
</P>

<P class="noindent" align="center">
<a name="img0274"></a>
<img src="images/0274.jpg" alt="[Illustration: And then she danced, verso]">
<img src="images/0275.jpg" alt="[Illustration: And then she danced, recto]">
</P>

<P>
And then she danced. . . .
</P>

<P>
<i>Where are you, Hyacinth?  There is a lover waiting for you somewhere,
my dear.</i>
</P>

<P>
It is the first of Spring.  The blackbird opens his yellow beak, and
whistles cool and clear.  There is blue magic in the morning; the sky,
deep-blue above, melts into white where it meets the hills.  The wind
waits for you up yonder&mdash;will you go to meet it?  Ah, stay here!  The
hedges have put on their green coats for you; misty green are the tall
elms from which the rooks are chattering.  Along the clean white road,
between the primrose banks, he comes. Will you be round this
corner?&mdash;&mdash;or the next?  He is looking for you, Hyacinth.
</P>

<P>
(She rested, breathless, and then danced again.)
</P>

<P>
It is summer afternoon.  All the village is at rest save one.
"Cuck-oo!" comes from the deep dark trees; "Cuck-oo!" he calls again,
and flies away to send back the answer.  The fields, all green and
gold, sleep undisturbed by the full river which creeps along them.
The air is heavy with the scent of may.  Where are you, Hyacinth?  Is
not this the trysting-place?  I have waited for you so long! . . .
</P>

<P>
She stopped, and the watcher in the bushes moved silently away, his
mind aflame with fancies.
</P>

<P>
Wiggs went back to the Palace to tell everybody that she could dance.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
"Shall we tell her how it happened?" said Udo jauntily.  "I just
recited a couple of lines&mdash;poetry, you know&mdash;backwards, and&mdash;well,
here I am!"
</P>

<P>
"O&mdash;&mdash;oh!" said Wiggs.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap16"></A>
<img src="images/0279X.jpg" alt="[Illustration: Detail of Belvane in an elaborate gown]">
</p>

<H3 ALIGN="center">
CHAPTER XVI
</H3>

<H3 ALIGN="center">
BELVANE ENJOYS HERSELF
</H3>

<P>
The entrance of an attendant into his room that morning to bring him
his early bran-mash had awakened Udo.  As soon as she was gone he
jumped up, shook the straw from himself, and said in a very passion of
longing,
</P>

<P class="poem">
     <i>Bo, boll, bill, bole.</i> <BR>
     <i>Wo, woll, will, wole.</i> <BR>
</P>

<P>
He felt it was his last chance.  Exhausted by his effort, he fell back
on the straw and dropped asleep again.  It was nearly an hour later
that he became properly awake.
</P>

<P>
Into his feelings I shall not enter at any length; I leave that to
Roger Scurvilegs.  Between ourselves Roger is a bit of a snob.  The
degradation to a Prince of Araby to be turned into an animal so
ludicrous, the delight of a Prince of Araby at regaining his own form,
it is this that he chiefly dwells upon.  Really, I think you or I
would have been equally delighted.  I am sure we can guess how Udo
felt about it.
</P>

<P>
He strutted about the room, he gazed at himself in every glass, he
held out his hand to an imaginary Hyacinth with "Ah, dear Princess,
and how are we this morning?"  Never had he felt so handsome and so
sure of himself.  It was in the middle of one of his pirouettings,
that he caught sight of the unfortunate bran-mash, and uttered the
remarkable words which I have already recorded.
</P>

<P>
The actual meeting with Hyacinth was even better than he had expected.
Hardly able to believe that it was true, she seized his hands
impulsively and cried:
</P>

<P>
"Oh, Prince Udo! oh, my dear, I <i>am</i> so glad!"
</P>

<P>
Udo twirled his moustache and felt a very gay dog indeed.
</P>

<P>
At breakfast  (where Udo did himself extremely well) they discussed
plans.  The first thing was to summon the Countess into their
presence.  An attendant was sent to fetch her.
</P>

<P>
"If you would like me to conduct the interview," said Udo, "I've no
doubt that&mdash;&mdash;"
</P>

<P>
"I think I shall be all right now that you are with me.  I shan't feel
so afraid of her now."
</P>

<P>
The attendant came in again.
</P>

<P>
"Her ladyship is not yet down, your Royal Highness."
</P>

<P>
"Tell her that I wish to see her directly she <i>is</i> down," said the
Princess.
</P>

<P>
The attendant withdrew.
</P>

<P>
"You were telling me about this army of hers," said Udo.  "One of my
ideas&mdash;I had a good many while I was&mdash;er&mdash;in retirement&mdash;was that she
could establish the army properly at her own expense, and that she
herself should be perpetual orderly-sergeant."
</P>

<P>
"Isn't that a nice thing to be?" asked Hyacinth innocently.
</P>

<P>
"It's a <i>horrible</i> thing to be.  Another of my ideas was that&mdash;&mdash;"
</P>

<P>
The attendant came in again.
</P>

<P>
"Her ladyship is a little indisposed, and is staying in bed for the
present."
</P>

<P>
"Oh!  Did her ladyship say when she thought of getting up?"
</P>

<P>
"Her ladyship didn't seem to think of getting up at all to-day.  Her
ladyship told me to say that she didn't seem to know <i>when</i> she'd get
up again."
</P>

<P>
The attendant withdrew, and Hyacinth and Udo, standing together in a
corner, discussed the matter anxiously.
</P>

<P>
"I don't quite see what we can <i>do</i>," said Hyacinth.  "We can't <i>pull</i>
her out of bed.  Besides, she may really be ill.  Supposing she stays
there for ever!"
</P>

<P>
"Of course," said Udo.  "It would be rather&mdash;&mdash;"
</P>

<P>
"You see if we&mdash;&mdash;"
</P>

<P>
"We might possibly&mdash;&mdash;"
</P>

<P>
"<i>Good</i> morning, all!" said Belvane, sweeping into the room.  She
dropped a profound curtsey to the Princess.  "Your Royal Highness!
And dear Prince Udo, looking his own charming self again!"
</P>

<P>
She had made a superb toilet.  In her flowing gold brocade, cut square
in front to reveal the whitest of necks, with her black hair falling
in two braids to her knees and twined with pearls which were caught up
in loops at her waist, she looked indeed a Queen; while Hyacinth and
Udo, taken utterly by surprise, seemed to be two conspirators whom she
had caught in the act of plotting against her.
</P>

<P class="noindent" align="center">
<a name="img0284"></a>
<img src="images/0284.jpg"
alt="[Illustration: &quot;Good morning,&quot; said Belvane, verso]">
<img src="images/0285.jpg"
alt="[Illustration: &quot;Good morning,&quot; said Belvane, recto]">
</P>

<P>
"I&mdash;I thought you weren't well, Countess," said Hyacinth, trying to
recover herself.
</P>

<P>
"I not well?" cried Belvane, clasping her hands to her breast.  "I
thought it was his Royal Highness who&mdash;&mdash;  Ah, but he's looking a true
Prince now."
</P>

<P>
She turned her eyes upon him, and there was in that look so much of
admiration, humour, appeal, impudence&mdash;I don't know what (and Roger
cannot tell us, either)&mdash;that Udo forgot entirely what he was going to
say and could only gaze at her in wonder.
</P>

<P>
Her mere entry dazzled him.  There is no knowing with a woman like
Belvane; and I believe she had purposely kept herself plain during
these last few days so that she might have the weapon of her beauty to
fall back upon in case anything went wrong.  Things had indeed gone
wrong; Udo had become a man again; and it was against the man that
this last weapon was directed.
</P>

<P>
Udo himself was only too ready.  The fact that he was once more
attractive to women meant as much as anything to him.  To have been
attractive to Hyacinth would have contented most of us, but Udo felt a
little uncomfortable with her.  He could not forget the last few days,
nor the fact that he had once been an object of pity to her.  Now
Belvane had not pitied him.
</P>

<P>
Hyacinth had got control of herself by this time.
</P>

<P>
"Enough of this, Countess," she said with dignity.  "We have not
forgotten the treason which you were plotting against the State; we
have not forgotten your base attack upon our guest, Prince Udo.  I
order you now to remain within the confines of the Palace until we
shall have decided what to do with you.  You may leave us."
</P>

<P>
Belvane dropped her eyes meekly.
</P>

<P>
"I am at your Royal Highness's commands.  I shall be in my garden when
your Royal Highness wants me."
</P>

<P>
She raised her eyes, gave one fleeting glance to Prince Udo, and
withdrew.
</P>

<P>
"A hateful woman," said Hyacinth.  "What shall we do with her?"
</P>

<P>
"I think," said Udo, "that I had better speak to her seriously first.
I have no doubt that I can drag from her the truth of her conspiracy
against you.  There may be others in it, in which case we shall have
to proceed with caution; on the other hand, it may be just misplaced
zeal on her part, in which case&mdash;&mdash;"
</P>

<P>
"Was it misplaced zeal which made her turn you into a&mdash;&mdash;?"
</P>

<P>
Udo held up his hand hastily.
</P>

<P>
"I have not forgotten that," he said.  "Be sure that I shall exact
full reparation.  Let me see; <i>which</i> is the way to her garden?"
</P>

<P>
Hyacinth did not know quite what to make of her guest.  At the moment
when she first saw him in his proper form the improvement on his late
appearance had been so marked that he had seemed almost the handsome
young Prince of her dreams.  Every minute after that had detracted
from him.  His face was too heavy, his manner was too pompous; one of
these days he would be too fat.
</P>

<P>
Moreover he was just a little too sure of his position in her house.
She had wanted his help, but she did not want so much of it as she
seemed to be likely to get.
</P>

<P>
Udo, feeling that it was going to be rather a nice day, went into
Belvane's garden.  He had been there once before; it seemed to him a
very much prettier garden this morning, and the woman who was again
awaiting him much more desirable.
</P>

<P>
Belvane made room for him on the seat next to her.
</P>

<P>
"This is where I sit when I write my poetry," she said.  "I don't know
if your Royal Highness is fond of poetry?"
</P>

<P>
"Extremely," said Udo.  "I have never actually written any or indeed
read much, but I have a great admiration for those who&mdash;er&mdash;admire it.
But it was not to talk about poetry that I came out here, Countess."
</P>

<P>
"No?" said Belvane.  "But your Royal Highness must have read the works
of Sacharino, the famous bard of Araby?"
</P>

<P>
"Sacharino, of course.  'Blood for something, something&mdash;&mdash;He who
something&mdash;&mdash;'  I mean, it's a delightful little thing.  Everybody
knows it.  But it was to talk about something very different that
I&mdash;&mdash;"
</P>

<P class="poem">
    "<i>Blood for blood and shoon for shoon,</i> <BR>
     &nbsp;<i>He who runs may read my rune,</i>"<BR>
</P>

<P>
quoted Belvane softly.  "It is perhaps Sacharino's most perfect gem."
</P>

<P>
"That's it," cried Udo excitedly.  "I knew I knew it, if only I
could&mdash;&mdash;"  He broke off suddenly, remembering the circumstances in
which he had wanted it.  He coughed importantly and explained for the
third time that he had not come to talk to her about poetry.
</P>

<P>
"But of course I think his most noble poem of all," went on Belvane,
apparently misunderstanding him, "is the ode to your Royal Highness
upon your coming-of-age.  Let me see, how does it begin?
</P>

<P class="poem">
    "<i>Prince Udo, so dashing and bold,</i> <BR>
     &nbsp;<i>Is apparently eighteen years old.</i> <BR>
         &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<i>It is eighteen years since</i><BR>
         &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<i>This wonderful Prince</i><BR>
     &nbsp;<i>Was born in the Palace, I'm told.</i>"<BR>
</P>

<P>
"These Court Poets," said Udo, with an air of unconcern, "flatter one,
of course."
</P>

<P>
If he expected a compliment he was disappointed.
</P>

<P>
"There I cannot judge," said Belvane, "until I know your Royal
Highness better."  She looked at him out of the corner of her eyes.
"Is your Royal Highness very&mdash;dashing?"
</P>

<P>
"I&mdash;er&mdash;well&mdash;er&mdash;one&mdash;that is to say."  He waded on uncomfortably,
feeling less dashing every moment.  He should have realised at once
that it was an impossible question to answer.
</P>

<P>
"Your Royal Highness," said Belvane modestly, "must not be too dashing
with us poor Euralians."
</P>

<P>
For the fourth time Udo explained that he had come there to speak to
her severely, and that Belvane seemed to have mistaken his purpose.
</P>

<P>
"Oh, forgive me, Prince Udo," she begged.  "I quite thought that you
had come out to commune soul to soul with a fellow-lover of the
beautiful."
</P>

<P>
"N&mdash;no," said Udo; "not exactly."
</P>

<P>
"Then what is it?" she cried, clasping her hands eagerly together.  "I
know it will be something exciting."
</P>

<P>
Udo stood up.  He felt that he could be more severe a little farther
off.  He moved a few yards away, and then turned round towards her,
resting his elbow on the sundial.
</P>

<P>
"Countess," he began sternly, "ten days ago, as I was starting on my
journey hither, I was suddenly&mdash;&mdash;"
</P>

<P>
"Just a moment," said Belvane, whispering eagerly to herself rather
than to him, and she jumped up with a cushion from the seat where she
was sitting, and ran across and arranged it under his elbow.  "He
would have been <i>so</i> uncomfortable," she murmured, and she hurried
back to her seat again and sat down and gazed at him, with her elbows
on her knees and her chin resting on her hands.  "Now go on telling
me," she said breathlessly.
</P>

<P>
Udo opened his mouth with the obvious intention of obeying her, but no
words came.  He seemed to have lost the thread of his argument.  He
felt a perfect fool, stuck up there with his elbow on a cushion, just
as if he were addressing a public meeting.  He looked at his elbow as
if he expected to find a glass of water there ready, and Belvane
divined his look and made a movement as if she were about to get it
for him.  It would be just like her.  He flung the cushion from him
("Oh, mind my roses," cried Belvane) and came down angrily to her.
Belvane looked at him with wide, innocent eyes.
</P>

<P>
"You&mdash;you&mdash;oh, <i>don't</i> look like that!"
</P>

<P>
"Like that?" said Belvane, looking like it again.
</P>

<P>
"Don't <i>do</i> it," shouted Udo, and he turned and kicked the cushion
down the flagged path.  "Stop it."
</P>

<P>
Belvane stopped it.
</P>

<P>
"Do you know," she said, "I'm rather frightened of you when you're
angry with me."
</P>

<P>
"I <i>am</i> angry.  Very, very angry.  Excessively annoyed."
</P>

<P>
"I thought you were," she sighed.
</P>

<P>
"And you know very well why."
</P>

<P>
She nodded her head at him.
</P>

<P>
"It's my dreadful temper," she said.  "I do such thoughtless things
when I lose my temper."
</P>

<P>
She sighed again and looked meekly at the ground.
</P>

<P>
"Er, well, you shouldn't," said Udo weakly.
</P>

<P>
"It was the slight to my sex that made me so angry.  I couldn't bear
to think that we women couldn't rule ourselves for such a short time,
and that a man had to be called in to help us."  She looked up at him
shyly.  "Of course I didn't know then what the man was going to be
like.  But now that I know&mdash;&mdash;"
</P>

<P>
Suddenly she held her arms out to him beseechingly.
</P>

<P>
"Stay with us, Prince Udo, and help us!  Men are so wise, so brave,
so&mdash;so generous.  They know nothing of the little petty feelings of
revenge that women indulge."
</P>

<P>
"Really, Countess, we&mdash;er&mdash;you&mdash;er&mdash;&mdash;  Of course there is a good deal
in what you say, and I&mdash;er&mdash;&mdash;"
</P>

<P>
"Won't you sit down again, Prince Udo?"
</P>

<P>
Udo sat down next to her.
</P>

<P>
"And now," said Belvane, "let's talk it over comfortably as friends
should."
</P>

<P>
"Of course," began Udo, "I quite see your point.  You hadn't seen me;
you didn't know anything about me; to you I might have been just any
man."
</P>

<P>
"I knew a little about you when you came here.  Beneath
the&mdash;er&mdash;outward mask I saw how brave and dignified you were.  But
even if I could have got you back into your proper form again, I think
I should have been afraid to; because I didn't know then how generous,
how forgiving you were."
</P>

<P>
It seemed to be quite decided that Udo was forgiving her.  When a very
beautiful woman thanks you humbly for something you have not yet given
her, there is only one thing for a gentleman to do.  Udo patted her
hand reassuringly.
</P>

<P>
"Oh, thank you, your Royal Highness."  She gave herself a little shake
and jumped up.  "And now shall I show you my beautiful garden?"
</P>

<P>
"A garden with you in it, dear Countess, is always beautiful," he said
gallantly.  And it was not bad, I think, for a man who had been living
on watercress and bran-mash only the day before.
</P>

<P>
They wandered round the garden together.  Udo was now quite certain it
was going to be a nice day.
</P>

<P>
It was an hour later when he came into the library. Hyacinth greeted
him eagerly.
</P>

<P>
"Well?" she said.
</P>

<P>
Udo nodded his head wisely.
</P>

<P>
"I have spoken to her about her conduct to me," he said.  "There will
be no more trouble in that direction, I fancy.  She explained her
conduct  to me very fully, and I have decided to overlook it this
time."
</P>

<P>
"But her robberies, her plots, her conspiracy against <i>me!</i>"
</P>

<P>
Udo looked blankly at her for a moment and then pulled himself
together.
</P>

<P>
"I am speaking to her about that this afternoon," he said.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap17"></A>
<img src="images/0299X.jpg" alt="[Illustration: Detail of the King of Barodia]">
</p>

<H3 ALIGN="center">
CHAPTER XVII
</H3>

<H3 ALIGN="center">
THE KING OF BARODIA DROPS THE WHISKER HABIT
</H3>

<P>
King Merriwig sat in his tent, his head held well back, his eyes
gazing upwards.  His rubicund cheeks were for the moment a snowy
white.  A hind of the name of Carlo had him firmly by the nose.  Yet
King Merriwig neither struggled nor protested; he was, in fact, being
shaved.
</P>

<P>
The Court Barber was in his usual conversational mood.  He released
his Majesty's nose for a moment, and, as he turned to sharpen his
razor, remarked,
</P>

<P>
"Terrible war, this."
</P>

<P>
"Terrible," agreed the King.
</P>

<P>
"Don't seem no end to it, like."
</P>

<P>
"Well, well," said Merriwig, "we shall see."
</P>

<P>
The barber got to work again.
</P>

<P>
"Do you know what I should do to the King of Barodia if I had him
here?"
</P>

<P>
Merriwig did not dare to speak, but he indicated with his right eye
that he was interested in the conversation.
</P>

<P>
"I'd shave his whiskers off," said Carlo firmly.
</P>

<P>
The King gave a sudden jerk, and for the moment there were signs of a
battle upon the snow; then the King leant back again, and in another
minute or so the operation was over.
</P>

<P>
"It will soon be all right," said Carlo, mopping at his Majesty's
chin.  "Your Majesty shouldn't have moved."
</P>

<P>
"It was my own fault, Carlo; you gave me a sudden idea, that's all."
</P>

<P>
"You're welcome, your Majesty."
</P>

<P>
As soon as he was alone the King took out his tablets.  On these he
was accustomed to record any great thoughts which occurred to him
during the day.  He now wrote in them these noble words:
</P>

<P>
"<i>Jewels of wisdom may fall from the meanest of hinds.</i>"
</P>

<P>
He struck a gong to summon the Chancellor into his presence.
</P>

<P>
"I have a great idea," he told the Chancellor.
</P>

<P>
The Chancellor hid his surprise and expressed his pleasure.
</P>

<P>
"To-night I propose to pay a secret visit to his Majesty the King of
Barodia.  Which of the many tents yonder have my spies located as the
royal one?"
</P>

<P>
"The big on in the centre, above which the Royal Arms fly."
</P>

<P>
"I thought as much.  Indeed I have often seen his Majesty entering it.
But one prefers to do these things according to custom.  Acting on
the information given me by my trusty spies, I propose to enter the
King of Barodia's tent at the dead of night, and&mdash;&mdash;"
</P>

<P>
The Chancellor shuddered in anticipation.
</P>

<P>
"And shave his whiskers off."
</P>

<P>
The Chancellor trembled with delight.
</P>

<P>
"Your Majesty," he said in a quavering voice, "forty years, man and
boy, have I served your Majesty, and your Majesty's late lamented
father, and never have I heard such a beautiful plan."
</P>

<P>
Merriwig struggled with himself for a moment, but his natural honesty
was too much for him.
</P>

<P>
"It was put into my head by a remark of my Court Barber's," he said
casually.  "But of course the actual working out of it has been mine."
</P>

<P>
"Jewels of wisdom," said the Chancellor sententiously, "may fall from
the meanest of hinds."
</P>

<P>
"I suppose," said Merriwig, taking up his tablets and absently
scratching out the words written thereon, "there is nothing in the
rules against it?"
</P>

<P>
"By no means, your Majesty.  In the annuals of Euralia there are many
instances of humour similar to that which your Majesty suggests:
humour, if I may say so, which, while evidencing to the ignorant only
the lighter side of war, has its roots in the most fundamental
strategical considerations."
</P>

<P>
Merriwig regarded him with admiration.  This was indeed a Chancellor.
</P>

<P>
"The very words," he answered, "which I said to myself when the idea
came to me.  'The fact,' I said, 'that this will help us to win the
war, must not disguise from us the fact that the King of Barodia will
look extremely funny without his whiskers.'  To-night I shall sally
forth and put my plan into practice."
</P>

<P>
At midnight, then, he started out.  The Chancellor awaited his return
with some anxiety.  This might well turn out to be the decisive stroke
(or strokes) of the war.  For centuries past the ruling monarchs of
Barodia had been famous for their ginger whiskers.  "As lost as the
King of Barodia without his whiskers" was indeed a proverb of those
times.  A King without a pair, and at such a crisis in his country's
fortunes!  It was inconceivable.  At the least he would have to live
in retirement until they grew again, and without the leadership of
their King the Barodian army would become a rabble.
</P>

<P>
The Chancellor was not distressed at the thought; he was looking
forward to his return to Euralia, where he kept a comfortable house.
It was not that his life in the field was uninteresting; he had as
much work to do as any man.  It was part of his business, for
instance, to test the pretentions of any new wizard or spell-monger
who was brought into the camp.  Such and such a quack would seek an
interview on the pretext that for five hundred crowns he could turn
the King of Barodia into a small black pig.  He would be brought
before the Chancellor.
</P>

<P>
"You say that you can turn a man into a small black pig?" the
Chancellor would ask.
</P>

<P>
"Yes, your lordship.  It came to me from my grandmother."
</P>

<P>
"Then turn me," the Chancellor would say simply.
</P>

<P>
The so-called wizard would try.  As soon as the incantation was over,
the Chancellor surveyed himself in the mirror. Then he nodded to a
couple of soldiers, and the impostor was tied backwards on to a mule
and driven with jeers out of the camp.  There were many such impostors
(who at least made a mule out of it), and the Chancellor's life did
not lack excitement.
</P>

<P>
But he yearned now for the simple comforts of his home.  He liked
pottering about his garden, when his work at the Palace was finished;
he liked, over the last meal of the day, to tell his wife all the
important things he had been doing since he had seen her, and to
impress her with the fact that he was the holder of many state secrets
which she must not attempt to drag from him.  A woman of less tact
would have considered the subject closed at this point, but she knew
that he was only longing to be persuaded.  However, as she always
found the secrets too dull to tell any one else, no great harm was
done.
</P>

<P>
"Just help me off with this cloak," said a voice in front of him.
</P>

<P>
The Chancellor felt about until his hands encountered a solid body.
He undid the cloak and the King stood revealed before him.
</P>

<P>
"Thanks.  Well, I've done it.  It went to my heart to do it at the
last moment, so beautiful they were, but I nerved myself to it.  Poor
soul, he slept like a lamb through it all.  I wonder what he'll say
when he wakes up."
</P>

<P>
"Did you bring them back with you?" asked the Chancellor excitedly.
</P>

<P>
"My dear Chancellor, what a question!"  He produced them from his
pocket.  "In the morning we'll run them up on the flagstaff for all
Barodia to see."
</P>

<P>
"He won't like that," said the Chancellor, chuckling.
</P>

<P>
"I don't quite see what he can do about it," said Merriwig.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
The King of Barodia didn't quite see either.
</P>

<P>
A fit of sneezing woke him up that morning, and at the same moment he
felt a curious draught about his cheeks.  He put his hand up and
immediately knew the worst.
</P>

<P>
"Hullo, there!" he bellowed to the sentry outside the door.
</P>

<P>
"Your Majesty," said the sentry, coming in with alacrity.
</P>

<P class="noindent" align="center">
<a name="img0308"></a>
<img src="images/0308.jpg" alt="[Illustration: The tent seemed to swim before his eyes, and he knew
 no more, verso]">
<img src="images/0309.jpg" alt="[Illustration: The tent seemed to swim before his eyes, and he knew
 no more, recto]">
</P>

<P>
The King bobbed down again at once.
</P>

<P>
"Send the Chancellor to me," said an angry voice from under the
bedclothes.
</P>

<P>
When the Chancellor came in it was to see the back only of his august
monarch.
</P>

<P>
"Chancellor," said the King, "prepare yourself for a shock."
</P>

<P>
"Yes, sir," said the Chancellor, trembling exceedingly.
</P>

<P>
"You are about to see something which no man in the history of Barodia
has ever seen before."
</P>

<P>
The Chancellor, not having the least idea what to expect, waited
nervously.  The next moment the tent seemed to swim before his eyes,
and he knew no more. . . .
</P>

<P>
When he came to, the King was pouring a jug of water down his neck and
murmuring rough words of comfort in his ear.
</P>

<P>
"Oh, your Majesty," said the poor Chancellor, "your Majesty!  I don't
know what to say, your Majesty."  He mopped at himself as he spoke,
and the water trickled from him on to the floor.
</P>

<P>
"Pull yourself together," said the King sternly.  "We shall want all
your wisdom, which is notoriously not much, to help us in this
crisis."
</P>

<P>
"Your Majesty, who has dared to do this grievous thing?"
</P>

<P>
"You fool, how should I know?  Do you think they did it while I was
awake?"
</P>

<P>
The Chancellor stiffened a little.  He was accustomed to being called
a fool; but that was by a man with a terrifying pair of ginger
whiskers.  From the rather fat and uninspiring face in front of him he
was inclined to resent it.
</P>

<P>
"What does your Majesty propose to do?" he asked shortly.
</P>

<P>
"I propose to do the following.  Upon you rests the chief burden."
</P>

<P>
The Chancellor did not look surprised.
</P>

<P>
"It will be your part to break the news as gently as possible to my
people.  You will begin by saying that I am busy with a great
enchanter who has called to see me, and that therefore I am unable to
show myself to my people this morning.  Later on in the day you will
announce that the enchanter has shown me how to defeat the wicked
Euralians; you will dwell upon the fact that this victory, as assured
by him, involves an overwhelming sacrifice on my part, but that for
the good of my people I am willing to endure it. Then you will
solemnly announce that the sacrifice I am making, have indeed already
made, is nothing less than&mdash;&mdash;  What are all those fools cheering for
out there?"  A mighty roar of laughter rose to the sky.  "Here, what's
it all about?  Just go and look."
</P>

<P>
The Chancellor went to the door of the tent&mdash;and saw.
</P>

<P>
He came back to the King, striving to speak casually.
</P>

<P>
"Just a humorous emblem that the Euralians have raised over their
camp," he said.  "It wouldn't amuse your Majesty."
</P>

<P>
"I am hardly in a mood for joking," said the King.  "Let us return to
business.  As I was saying, you will announce to the people that the
enormous sacrifice which their King is prepared to make for them
consists of&mdash;  There they go again.  I must really see what it is.
Just pull the door back so that I may see without being seen."
</P>

<P>
"It&mdash;it really wouldn't amuse your Majesty."
</P>

<P>
"Are you implying that I have no sense of humour?" said the King
sternly.
</P>

<P>
"Oh, no, sire, but there are certain jokes, jokes in the poorest of
taste, that would naturally not appeal to so delicate a palate as your
Majesty's.  This&mdash;er&mdash;strikes me as one of them."
</P>

<P>
"Of that I am the best judge," said the King coldly.  "Open the door
at once."
</P>

<P>
The Chancellor opened the door; and there before the King's eyes,
flaunting themselves in the breeze beneath the Royal Standard of
Euralia, waved his own beloved whiskers.
</P>

<P>
The King of Barodia was not a lovable man, and his daughters were
decidedly plain, but there are moments when one cannot help admiring
him.  This was one of them.
</P>

<P>
"You may shut the door," he said to the Chancellor.  "The instructions
which I gave to you just now," he went on in the same cold voice, "are
cancelled.  Let me think for a moment."  He began to walk up and down
his apartment.  "You may think, too," he added kindly.  "If you have
anything not entirely senseless to suggest, you may suggest it."
</P>

<P>
He continued his pacings.  Suddenly he came to a dead stop.  He was
standing in front of a large mirror.  For the first time since he was
seventeen he had seen his face without whiskers.  His eyes still fixed
on his reflection, he beckoned the Chancellor to approach.
</P>

<P>
"Come here," he said, clutching him by the arm.  "You see that?"  He
pointed to the reflection.  "That is what I look like?  The mirror
hasn't made a mistake of any kind?  That is really and truly what I
look like?"
</P>

<P>
"Yes, sire."
</P>

<P>
For a little while the King continued to gaze fascinated at his
reflection, and then he turned on the Chancellor.
</P>

<P>
"You coward!" he said.  "You weak-kneed, jelly-souled, paper-livered
imitation of a man!  You cringe to a King who looks like that!  Why,
you ought to <i>kick</i> me."
</P>

<P>
The Chancellor remembered that he had one kick owing to him.  He drew
back his foot, and then a thought occurred to him.
</P>

<P>
"You might kick me back," he pointed out.
</P>

<P>
"I certainly should," said the King.
</P>

<P>
The Chancellor hesitated a moment.
</P>

<P>
"I think," he said, "that these private quarrels in the face of the
common enemy are to be deplored."
</P>

<P>
The King looked at him, gave a short laugh, and went on walking up and
down.
</P>

<P>
"That face again," he sighed as he came opposite the mirror.  "No,
it's no good; I can never be King like this.  I shall abdicate."
</P>

<P>
"But, your Majesty, this is a very terrible decision.  Could not your
Majesty live in retirement until your Majesty had grown your Majesty's
whiskers again?  Surely this is&mdash;&mdash;"
</P>

<P>
The King came to a stand opposite him and looked down on him gravely.
</P>

<P>
"Chancellor," he said, "those whiskers which you have just seen
fluttering in the breeze have been for more than forty years my curse.
For more than forty years I have had to live up to those whiskers,
behaving, not as my temperament, which is a kindly, indeed a genial
one, bade me to behave, but as those whiskers insisted I should
behave.  Arrogant, hasty-tempered, over-bearing&mdash;these are the
qualities which have been demanded of the owner of those whiskers.  I
played a part which was difficult at first; of late, it has, alas!
been more easy.  Yet it has never been my true nature that you have
seen."
</P>

<P>
He paused and looked silently at himself in the glass.
</P>

<P>
"But, your Majesty," said the Chancellor eagerly, "why choose this
moment to abdicate?  Think how your country will welcome this new King
whom you have just revealed to me.  And yet," he added regretfully,
"it would not be quite the same."
</P>

<P>
The King turned round to him.
</P>

<P>
"There spoke a true Barodian," he said.  "It would not be the same.
Barodians have come to expect certain qualities from their rulers, and
they would be lost without them.  A new King might accustom them to
other ways, but they are used to me, and they would not like me
different.  No, Chancellor, I shall abdicate.  Do not wear so sad a
face for me.  I am looking forward to my new life with the greatest of
joy."
</P>

<P>
The Chancellor was not looking sad for him; he was looking sad for
himself, thinking that perhaps a new King might like changes in
Chancellors equally with changes in manners or whiskers.
</P>

<P>
"But what will you do?" he asked.
</P>

<P>
"I shall be a simple subject of the new King, earning my living by my
own toil."
</P>

<P>
The Chancellor raised his eyebrows at this.
</P>

<P>
"I suppose you think," said the King haughtily, "that I have not the
intelligence to earn my own living."
</P>

<P>
The Chancellor with a cough remarked that the very distinguished
qualities which made an excellent King did not always imply the
corresponding&mdash;er&mdash;and so on.
</P>

<P>
"That shows how little you know about it.  Just to give one example.
I happen to know that I have in me the makings of an excellent
swineherd."
</P>

<P>
"A swineherd?"
</P>

<P>
"The man who&mdash;er&mdash;herds the swine.  It may surprise you to hear that,
posing as a swineherd, I have conversed with another of the profession
upon his own subject, without his suspecting the truth.  It is just
such a busy outdoor life as I should enjoy.  One herds and one milks,
and one milks, and&mdash;er&mdash;herds, and so it goes on day after day."  A
happy smile, the first the Chancellor had ever seen there, spread
itself over his features.  He clapped the Chancellor playfully on the
back and added, "I shall simply love it."
</P>

<P>
The Chancellor was amazed.  What a story for his dinner-parties when
the war was over!
</P>

<P>
"How will you announce it?" he asked, and his tone struck a happy mean
between the tones in which you address a monarch and a pig-minder
respectively.
</P>

<P>
"That will be your duty.  Now that I have shaken off the curse of
those whiskers, I am no longer a proud man, but even a swineherd would
not care for it to get about that he had been forcibly shaved while
sleeping.  That this should be the last incident recorded of me in
Barodian history is unbearable.  You will announce therefore that I
have been slain in fair combat, though at the dead of night, by the
King of Euralia, and that my whiskers fly over his royal tent as a
symbol of his victory."  He winked at the Chancellor and added, "It
might as well get about that some one had stolen my Magic Sword that
evening."
</P>

<P>
The Chancellor was speechless with admiration and approval of the
plan.  Like his brother of Euralia, he too was longing to get home
again.  The war had arisen over a personal insult to the King.  If the
King was no longer King, why should the war go on?
</P>

<P>
"I think," said the future swineherd, "that I shall send a Note over
to the King of Euralia, telling him my decision.  To-night, when it is
dark, I shall steal away and begin my new life.  There seems to be no
reason why the people should not go back to their homes to-morrow.  By
the way, that guard outside there knows that I wasn't killed last
night; that's rather awkward."
</P>

<P>
"I think," said the Chancellor, who was already picturing his return
home, and was not going to be done out of it by a common sentry, "I
think I could persuade him that you <i>were</i> killed last night."
</P>

<P>
"Oh, well, then, that's all right."  He drew a ring from his finger.
"Perhaps this will help him to be persuaded. Now leave me while I
write to the King of Euralia."
</P>

<P>
It was a letter which Merriwig was decidedly glad to get. It announced
bluntly that the war was over, and added that the King of Barodia
proposed to abdicate.  His son would rule in his stead, but he was a
harmless fool, and the King of Euralia need not bother about him.  The
King would be much obliged if he would let it get about that the
whiskers had been won in a fair fight; this would really be more to
the credit of both of them.  Personally he was glad to be rid of the
things, but one has one's dignity.  He was now retiring into private
life, and if it were rumoured abroad that he had been killed by the
King of Euralia matters would be much more easy to arrange.
</P>

<P>
Merriwig slept late after his long night abroad, and he found this
Note waiting for him when he awoke.  He summoned the Chancellor at
once.
</P>

<P>
"What have you done about those&mdash;er&mdash;trophies?" he asked.
</P>

<P>
"They are fluttering from your flagstaff, sire, at this moment."
</P>

<P>
"Ah!  And what do my people say?"
</P>

<P>
"They are roaring with laughter, sire, at the whimsical nature of the
jest."
</P>

<P>
"Yes, but what do they say?"
</P>

<P>
"Some say that your Majesty, with great cunning, ventured privily in
the night and cut them off while he slept; others, that with great
bravery you defeated him in mortal combat and carried them away as the
spoils of the victor."
</P>

<P>
"Oh!  And what did <i>you</i> say?"
</P>

<P>
The Chancellor looked reproachful.
</P>

<P>
"Naturally, your Majesty, I have not spoken with them."
</P>

<P>
"Ah, well, I have been thinking it over in the night, and I remember
now that I <i>did</i> kill him.  You understand?"
</P>

<P>
"Your Majesty's skill in sword play will be much appreciated by the
people."
</P>

<P>
"Quite so," said the King hastily.  "Well, that's all&mdash;I'm getting up
now.  And we're all going home to-morrow."
</P>

<P>
The Chancellor went out, rubbing his hands with delight.
</P>

<p class="noindent" align="right">
<img src="images/0323X.jpg" alt="[Illustration: Small picture of a thin man carrying a large sack]">
</p>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap18"></A>
<img src="images/0325X.jpg" alt="[Illustration: A small girl in medieval garb holds a large document">
</p>

<H3 ALIGN="center">
CHAPTER XVIII
</H3>

<H3 ALIGN="center">
THE VETERAN OF THE FOREST ENTERTAINS TWO VERY YOUNG PEOPLE
</H3>

<P>
Do you remember the day when the Princess Hyacinth and Wiggs sat upon
the castle walls and talked of Udo's coming?  The Princess thought he
would be dark, and Wiggs thought he would be fair, and he was to have
the Purple Room&mdash;or was it the Blue?&mdash;and anyhow he was to put the
Countess in her place and bring happiness to Euralia.  That seemed a
long time ago to Hyacinth now, as once more she sat on the castle
walls with Wiggs.
</P>

<P>
She was very lovely.  She longed to get rid of that "outside help in
our affairs" which she had summoned so recklessly.  They were two
against one now.  Belvane actively against her was bad enough; but
Belvane in the background with Udo as her mouthpiece&mdash;Udo specially
asked in to give the benefit of his counsel&mdash;this was ten times worse.
</P>

<P>
"What do you do, Wiggs?" she asked, "when you are very lonely and
nobody loves you?"
</P>

<P>
"Dance," said Wiggs promptly.
</P>

<P>
"But if you don't want to dance?"
</P>

<P>
Wiggs tried to remember those dark ages (about a week ago) when she
couldn't dance.
</P>

<P>
"I used to go into the forest," she said, "and sit under my own tree,
and by and by everybody loved you."
</P>

<P>
"I wonder if they'd love <i>me</i>."
</P>

<P>
"Of course they would.  Shall I show you my special tree?"
</P>

<P>
"Yes, but don't come with me; tell me where it is.  I want to be
unhappy alone."
</P>

<P>
So Wiggs told her how you followed her special path, which went in at
the corner of the forest, until by and by the trees thinned on either
side, and it widened into a glade, and you went downhill and crossed
the brook at the bottom and went up the other side until it was all
trees again, and the first and the biggest and the oldest and the
loveliest was hers.  And you turned round and sat with your back
against it, and looked across to where you'd come from, and then you
<i>knew</i> that everything was all right.
</P>

<P>
"I shall find it," said Hyacinth, as she got up.  "Thank you, dear."
</P>

<P>
She found it, she sat there, and her heart was very bitter at first
against Udo and against Belvane, and even against her father for going
away and leaving her; but by and by the peace of the place wrapped
itself around her, and she felt that she would find a way out of her
difficulties somehow.  Only she wished that her father would come
back, because he loved her, and she felt that it would be nice to be
loved again.
</P>

<P>
"It is beautiful, isn't it?" said a voice from behind her.
</P>

<P>
She turned suddenly, as a tall young man stepped out from among the
trees.
</P>

<P>
"Oh, who are you, please?" she asked, amazed at his sudden appearance.
His dress told her nothing, but his face told her things which she
was glad to know.
</P>

<P>
"My name," he said, "is Coronel."
</P>

<P>
"It is a pretty name."
</P>

<P>
"Yes, but don't be lead away by it.  It belongs to nobody very
particular.  Do you mind if I sit down?  I generally sit down here
about this time."
</P>

<P>
"Oh, do you live in the forest?"
</P>

<P>
"I have lived here for the last week."  He gave her a friendly smile,
and added, "You're late, aren't you?"
</P>

<P>
"Late?"
</P>

<P>
"Yes, I've been expecting you for the last seven days."
</P>

<P>
"How did you know there was any me at all?" smiled Hyacinth.
</P>

<P>
With a movement of his hand Coronel indicated the scene in front of
him.
</P>

<P>
"There had to be <i>somebody</i> for whom all this was made.  It wanted
somebody to say thank you to it now and then."
</P>

<P>
"Haven't you been doing that all this week?"
</P>

<P>
"Me?  I wouldn't presume.  No, it's your glade, and you've neglected
it shamefully."
</P>

<P>
"There's a little girl who comes here," said Hyacinth.  "I wonder if
you have seen her?"
</P>

<P>
Coronel turned away.  There were secret places in his heart into which
Hyacinth could not come&mdash;yet.
</P>

<P>
"She danced," he said shortly.
</P>

<P>
There was silence between them for a little, but a comfortable
silence, as if they were already old friends.
</P>

<P>
"You know," said Hyacinth, looking down at him as he lay at her feet,
"you ought not to be here at all, really."
</P>

<P>
"I wish I could think that," said Coronel.  "I had a horrible feeling
that duty called me here.  I love those places where one really
oughtn't to be at all, don't you?"
</P>

<P>
"I love being here," sighed Hyacinth.  "Wiggs was quite right."
Seeing him look up at her she added, "Wiggs is the little girl who
dances, you know."
</P>

<P>
"She would be right," said Coronel, looking away from her.
</P>

<P>
Hyacinth felt strangely rested.  It seemed that never again would
anything trouble her; never again would she have only her own strength
to depend upon.  Who was he?  But it did not matter.  He might go away
and she might never see him again, but she was no longer afraid of the
world.
</P>

<P>
"I thought," she said, "that all the men of Euralia were away
fighting."
</P>

<P>
"So did I," said Coronel.
</P>

<P>
"What are you, then?  A Prince from a distant country, an enchanter, a
spy sent from Barodia, a travelling musician?&mdash;you see, I give you
much to choose from."
</P>

<P>
"You leave me nothing to be but what I am&mdash;Coronel."
</P>

<P>
"And I am Hyacinth."
</P>

<P>
He knew, of course, but he made no sign.
</P>

<P>
"Hyacinth," he said, and he held out his hand.
</P>

<P>
"Coronel," she answered as she took it.
</P>

<P>
The brook chuckled to itself as it hurried past below them.
</P>

<P>
Hyacinth got up with a little sigh of contentment.
</P>

<P>
"Well, I must be going," she said.
</P>

<P>
"Must you really be going?" asked Coronel.  "I wasn't saying good-bye,
you know."
</P>

<P class="noindent" align="center">
<a name="img0332"></a>
<img src="images/0332.jpg"
alt="[Illustration: She turned round and went off daintily down the hill, verso]">
<img src="images/0333.jpg"
alt="[Illustration: She turned round and went off daintily down the hill, verso]">
</P>

<P>
"I really must."
</P>

<P>
"It's a surprising thing about the view from here," said Coronel,
"that it looks just as nice to-morrow.  To-morrow about the same
time."
</P>

<P>
"That's a very extraordinary thing," smiled Hyacinth.
</P>

<P>
"Yes, but it's one of those things that you don't want to take another
person's word for."
</P>

<P>
"You think I ought to see for myself?  Well, perhaps I will."
</P>

<P>
"Give me a whistle if I happen to be passing," said Coronel casually,
"and tell me what you think.  Good-bye, Hyacinth."
</P>

<P>
"Good-bye, Coronel."
</P>

<P>
She nodded her head confidently at him, and then turned round and went
off daintily down the hill.
</P>

<P>
Coronel stared after her.
</P>

<P>
"What <i>is</i> Udo doing?" he murmured to himself.  "But perhaps she
doesn't like animals.  A whole day to wait. How endless!"
</P>

<P>
If he had known that Udo, now on two legs again, was at that moment in
Belvane's garden, trying to tell her, for the fifth time that week,
about his early life in Araby, he would have been still more
surprised.
</P>

<P>
We left Coronel, if you remember, in Araby.  For three or four days he
remained there, wondering how Udo was getting on, and feeling more and
more that he ought to do something about it.  On the fourth day he got
on to his horse and rode off again.  He simply must see what was
happening.  If Udo wanted to help, then he would be there to give it;
if Udo was all right again, then he could go comfortably back to
Araby.
</P>

<P>
To tell the truth, Coronel was a little jealous of his friend.  A
certain Prince Perivale, who had stayed at his uncle's court, had once
been a suitor for Hyacinth's hand; but losing a competition with the
famous seven-headed bull of Euralia, which Merriwig had arranged for
him, had made no further headway with his suit.  This Prince had had a
portrait of Hyacinth specially done for him by his own Court Painter,
a portrait which Coronel had seen.  It was for this reason that he had
at first objected to accompanying Udo to Euralia, and it was for this
reason that he persuaded himself very readily that the claims of
friendship called him there now.
</P>

<P>
For the last week he had been waiting in the forest.  Now that he was
there, he was not quite sure how to carry out his mission.  So far
there had been no sign of Udo, either on four legs or on two; it
seemed probable that unless Coronel went to the Palace and asked for
him, there would be no sign.  And if he went to the Palace, and Udo
was all right, and the Princess Hyacinth was in love with him, then
the worst would have happened.  He would have to stay there and help
admire Udo&mdash;an unsatisfying prospect to a man in love.  For he told
himself by this time that he was in love with Hyacinth, although he
had never seen her.
</P>

<P>
So he had waited in the forest, hoping for something to turn up; and
first Wiggs had come . . . and now at last Hyacinth.  He was very glad
that he had waited.
</P>

<P>
She was there on the morrow.
</P>

<P>
"I knew you'd come," said Coronel.  "It looks just as beautiful,
doesn't it?"
</P>

<P>
"I think it's even more beautiful," said Hyacinth.
</P>

<P>
"You mean those little white clouds?  That was my idea putting those
in.  I thought you'd like them."
</P>

<P>
"I wondered what you did all day.  Does it keep you very busy?"
</P>

<P>
"Oh," said Coronel, "I have time for singing."
</P>

<P>
"Why do you sing?"
</P>

<P>
"Because I am young and the forest is beautiful."
</P>

<P>
"I have been singing this morning, too."
</P>

<P>
"Why?" asked Coronel eagerly.
</P>

<P>
"Because the war with Barodia is over."
</P>

<P>
"Oh!" said Coronel, rather taken aback.
</P>

<P>
"That doesn't interest you.  Yet if you were a Euralian&mdash;&mdash;"
</P>

<P>
"But it interests me extremely.  Let us admire the scene for a moment,
while I think.  Look, there is another of my little clouds."
</P>

<P>
Coronel wondered what would happen now.  If the King were coming back,
then Udo would be wanted no longer save as a suitor for Hyacinth's
hand.  If, then, he returned, it would show that&mdash;&mdash;  But suppose he
was still an animal? It was doubtful if he would go back to Araby as
an animal.  And then there was another possibility: perhaps he had
never come to Euralia at all.  Here were a lot of questions to be
answered, and here next to him was one who could answer them.  But he
must go carefully.
</P>

<P>
"Ninety-seven, ninety-eight, ninety-nine, a hundred," he said aloud.
"There, I've finished my thinking and you've finished your looking."
</P>

<P>
"And what have you decided?" smiled Hyacinth.
</P>

<P>
"Decided?" said Coronel, rather startled.  "Oh, no, I wasn't deciding
anything, I was just thinking.  I was thinking about animals."
</P>

<P>
"So was I."
</P>

<P>
"How very curious, and also how wrong of you.  You were supposed to be
admiring my clouds.  What sort of animals were you thinking about?"
</P>

<P>
"Oh&mdash;all sorts."
</P>

<P>
"I was thinking about rabbits.  Do you care for rabbits at all?"
</P>

<P>
"Not very much."
</P>

<P>
"Neither do I.  They're so loppity.  Do you like lions?"
</P>

<P>
"I think their tails are rather silly," said Hyacinth.
</P>

<P>
"Yes, perhaps they are.  Now&mdash;a woolly lamb."
</P>

<P>
"I am not very fond of woolly lambs just now."
</P>

<P>
"No?  Well, they're not very interesting.  It's a funny thing," he
went on casually, trying to steal a glance at her, "that we should be
talking about those three animals, because I once met somebody who was
a mixture of all three together at the same time."
</P>

<P>
"So did I," said Hyacinth gravely.
</P>

<P>
But he saw her mouth trembling, and suddenly she turned round and
caught his eye, and then they burst out laughing together.
</P>

<P>
"Poor Udo," said Coronel; "and how is he looking now?"
</P>

<P>
"He is all right again now."
</P>

<P>
"All right again?  Then why isn't he&mdash;&mdash;  But I'm very glad he isn't."
</P>

<P>
"I didn't like him," said Hyacinth, blushing a little.  And then she
went on bravely, "But I think he found he didn't like me first."
</P>

<P>
"He wants humouring," said Coronel.  "It's my business to humour him,
it isn't yours."
</P>

<P>
Hyacinth looked at him with a new interest.
</P>

<P>
"Now I know who you are," she said.  "He talked about you once."
</P>

<P>
"What did he say?" asked Coronel, obviously dying to know.
</P>

<P>
"He said you were good at poetry."
</P>

<P>
Coronel was a little disappointed.  He would have preferred Hyacinth
to have been told that he was good at dragons.  However, they had met
now and it did not matter.
</P>

<P>
"Princess," he said suddenly, "I expect you wonder what I am doing
here.  I came to see if Prince Udo was in need of help, and also to
see if you were in need of help.  Prince Udo was my friend, but if he
has not been a friend of yours, then he is no longer a friend of mine.
Tell me what has been happening here, and then tell me if in any way
I can help you."
</P>

<P>
"You called me Hyacinth yesterday," she said, "and it is still my
name."
</P>

<P>
"Hyacinth," said Coronel, taking her hand, "tell me if you want me at
all."
</P>

<P>
"Thank you, Coronel.  You see, Coronel, it's like this."  And sitting
beneath Wiggs's veteran of the forest, with Coronel lying at her feet,
she told him everything.
</P>

<P>
"It seems easy enough," he said when she had finished.  "You want Udo
pushed out and the Countess put in her place.  I can do the one while
you do the other."
</P>

<P>
"Yes, but how do I push Prince Udo out?"
</P>

<P>
"That's what <i>I'm</i> going to do."
</P>

<P>
"Yes, but, Coronel dear, if I could put the Countess in her place,
shouldn't I have done it a long time ago?  I don't think you quite
know the sort of person she is.  And I don't quite know what her place
is either, which makes it rather had to put her into it.  You see, I
don't think I told you that&mdash;that Father is rather fond of her."
</P>

<P>
"I thought you said Udo was."
</P>

<P>
"They both are."
</P>

<P>
"Then how simple.  We simply kill Udo, and&mdash;and&mdash;well, anyhow, there's
one part of it done."
</P>

<P>
"Yes, but what about the other part?"
</P>

<P>
Coronel thought for a moment.
</P>

<P>
"Would it be simpler if we did it the other way around?" he said.
"Killed the Countess and put Udo in his place."
</P>

<P>
"Father wouldn't like that at all, and he's coming back to-morrow."
</P>

<P>
Coronel didn't quite see the difficulty.  If the King was in love with
the Countess, he would marry her whatever Hyacinth did.  And what was
the good of putting her in her place for one day if her next place was
to be on the throne.
</P>

<P>
Hyacinth guessed what he was thinking.
</P>

<P>
"Oh, don't you see," she cried, "she doesn't know that the King is
coming back to-morrow.  And if I can only just show her&mdash;I don't mind
if it's only for an hour&mdash;that I am not afraid of her, and that she
has got to take her orders from me, then I shan't mind so much all
that has happened these last weeks.  But if she is to have disregarded
me all the time, if she is to have plotted against me from the very
moment my father went away, and if nothing is to come to her for it
but that she marries my father and becomes Queen of Euralia, then I
can have no pride left, and I will be a Princess no longer."
</P>

<P>
"I must see this Belvane," said Coronel thoughtfully.
</P>

<P>
"Oh, Coronel, Coronel," cried Hyacinth, "if <i>you</i> fall in love with
her, too, I think I shall die of shame!"
</P>

<P>
"With <i>her</i>, Hyacinth?" he said, turning to her in amazement.
</P>

<P>
"Yes, you&mdash;I didn't&mdash;you never&mdash;I&mdash;&mdash;"  Her voice trailed away; she
could not meet his gaze any longer; she dropped her eyes, and the next
moment his arms were round her, and she knew that she would never be
alone again.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap19"></A>
<img src="images/0347X.jpg" alt="[Illustration: Detail of Hyacinth presenting Coronel]">
</p>

<H3 ALIGN="center">
CHAPTER XIX
</H3>

<H3 ALIGN="center">
UDO BEHAVES LIKE A GENTLEMAN
</H3>

<P>
"And now," said Coronel, "we'd better decide what to do."
</P>

<P>
"But I don't mind what we do now," said Hyacinth happily.  "She may
have the throne and Father and Udo, and&mdash;and anything else she can
get, and I shan't mind a bit.  You see, I have got <i>you</i> now, Coronel,
and I can never be jealous of anybody again."
</P>

<P>
"That's what makes it so jolly.  We can do what we like, and it
doesn't matter if it doesn't come off.  So just for fun let's think of
something to pay her out."
</P>

<P>
"I feel I don't want to hurt anybody to-day."
</P>

<P>
"All right, we won't hurt her, we'll humour her.  We will be her most
humble obedient servants.  She shall have everything she wants."
</P>

<P>
"Including Prince Udo," smiled Hyacinth.
</P>

<P>
"That's a splendid idea.  We'll make her have Udo.  It will annoy your
father, but one can't please everybody.  Oh, I can see myself enjoying
this."
</P>

<P>
They got up and wandered back along Wiggs's path, hand in hand.
</P>

<P>
"I'm almost afraid to leave the forest," said Hyacinth, "in case
something happens."
</P>

<P>
"What should happen?"
</P>

<P>
"I don't know; but all our life together has been in the forest, and
I'm just a little afraid of the world."
</P>

<P>
"I will be very close to you always, Hyacinth."
</P>

<P>
"Be very close, Coronel," she whispered, and then they walked out
together.
</P>

<P>
If any of the servants at the Palace were surprised to see Coronel,
they did not show it.  After all, that was their business.
</P>

<P>
"Prince Coronel will be staying here," said the Princess.  "Prepare a
room for him and some refreshment for us both."  And if they discussed
those things in the servants' halls of those days (as why should they
not?), no doubt they told each other that the Princess Hyacinth (bless
her pretty face!) had found her man at last.  Why, you only had to see
her looking at him.  But I get no assistance from Roger at this point;
he pretends that he has a mind far above the gossip of the lower
orders.
</P>

<P>
"I say," said Coronel, as they went up the grand staircase, "I am not
a Prince, you know.  Don't say I have deceived you."
</P>

<P>
"You are <i>my</i> Prince," said Hyacinth proudly.
</P>

<P>
"My dear, I am a king among men to-day, and you are my queen, but
that's in our own special country of two."
</P>

<P>
"If you are so particular," said Hyacinth, with a smile, "Father will
make you a proper Prince directly he comes back."
</P>

<P>
"Will he?  That's what I'm wondering.  You see he doesn't know yet
about our little present to the Countess."
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
But it is quite time we got back to Belvane; we have left her alone
too long.  It was more than Udo did.  Just now he was with her in her
garden, telling her for the fifth time an extraordinarily dull story
about an encounter of his with a dragon, apparently in its dotage, to
which Belvane was listening with an interest which surprised even the
narrator.
</P>

<P>
"And then," said Udo, "I jumped quickly to the right, and whirling
my&mdash;no, wait a bit, that was later&mdash;I jumped quickly to my left&mdash;yes,
I remember it now, it <i>was</i> my left&mdash;I jumped quickly to my left, and
whirling my&mdash;&mdash;"
</P>

<P>
He stopped suddenly at the expression on Belvane's face.  She was
looking over his shoulder at something behind him.
</P>

<P>
"Why, whoever is this?" she said, getting to her feet.
</P>

<P>
Before Udo had completely cleared his mind of his dragon, the Princess
and Coronel were upon them.
</P>

<P>
"Ah, Countess, I thought we should find you together," said Hyacinth
archly.  "Let me present to you my friend, the Duke Coronel.  Coronel,
this is Countess Belvane, a very dear and faithful friend of mine.
Prince Udo, of course, you know. His Royal Highness and the Countess
are&mdash;well, it isn't generally known at present, so perhaps I oughtn't
to say anything."
</P>

<P>
Coronel made a deep bow to the astonished Belvane.
</P>

<P class="noindent" align="center">
<a name="img0352"></a>
<img src="images/0352.jpg"
alt="[Illustration: Let me present to you my friend the Duke Coronel, verso]">
<img src="images/0353.jpg"
alt="[Illustration: Let me present to you my friend the Duke Coronel, recto]">
</P>

<P>
"Your humble servant," he said.  "You will, I am sure, forgive me if I
say how glad I am to hear your news.  Udo is one of my oldest
friends"&mdash;he turned and clapped that bewildered Highness on the
back&mdash;"aren't you, Udo? and I can think of no one more suitable in
every way."  He bowed again, and turned back to the Prince.  "Well,
Udo, you're looking splendid.  A different thing, Countess, from when
I last saw him.  Let me see, that must have been just the day before
he arrived in Euralia.  Ah, what a miracle-worker True Love is!"
</P>

<P>
I think one of the things which made Belvane so remarkable was that
she was never afraid of remaining silent when she was not quite sure
what to say.  She waited therefore while she considered what all this
meant; who Coronel was, what he was doing there, even whether a
marriage with Udo was not after all the best that she could hope for
now.
</P>

<P>
Meanwhile Udo, of course, blundered along gaily.
</P>

<P>
"We aren't exactly, Princess&mdash;I mean&mdash;&mdash;What are you doing here,
Coronel?&mdash;I didn't know, Princess, that you&mdash;&mdash;  The Countess and I
were just having a little&mdash;I was just telling her what you said
about&mdash;How did you get here, Coronel?"
</P>

<P>
"Shall we tell him?" said Coronel, with a smile at Hyacinth.
</P>

<P>
Hyacinth nodded.
</P>

<P>
"I rode," said Coronel.  "It's a secret," he added.
</P>

<P>
"But I didn't know that you&mdash;&mdash;"
</P>

<P>
"We find that we have really known each other a very long time,"
explained Hyacinth.
</P>

<P>
"And hearing that there was to be a wedding," added Coronel&mdash;&mdash;
</P>

<P>
Belvane made up her mind.  Coronel was evidently a very different man
from Udo.  If he stayed in Euralia as adviser&mdash;more than adviser she
guessed&mdash;to Hyacinth, her own position would not be in much doubt.
And as for the King, it might be months before he came back, and when
he did come would he remember her?  But to be Queen of Araby was no
mean thing.
</P>

<P>
"We didn't want it to be known yet," she said shyly, "but you have
guessed our secret, your Royal Highness."  She looked modestly at the
ground, and, feeling for her reluctant lover's hand, went on, "Udo and
I"&mdash;here she squeezed the hand, and, finding it was Coronel's, took
Udo's boldly without any more maidenly nonsense&mdash;"Udo and I love each
other."
</P>

<P>
"Say something, Udo," prompted Coronel.
</P>

<P>
"Er&mdash;yes," said Udo, very unwillingly, and deciding he would explain
it all afterwards.  Whatever his feelings for the Countess, he was not
going to be rushed into a marriage.
</P>

<P>
"Oh, I'm so glad," said Hyacinth.  "I felt somehow that it must be
coming, because you've seen so <i>much</i> of each other lately.  Wiggs and
I have often talked about it together."
</P>

<P>
("What has happened to the child?" thought Belvane.  "She isn't a
child at all, she's grown up.")
</P>

<P>
"There's no holding Udo once he begins," volunteered Coronel.  "He's
the most desperate lover in Araby.
</P>

<P>
"My father will be so excited when he hears," said Hyacinth.  "You
know, of course, that his Majesty comes back to-morrow with all his
army."
</P>

<P>
She did not swoon or utter a cry.  She did not plead the vapours or
the megrims.  She took unflinching what must have been the biggest
shock in her life.
</P>

<P>
"Then perhaps I had better see that everything is ready in the
Palace," she said, "if your Royal Highness will excuse me."  And with
a curtsey she was gone.
</P>

<P>
Coronel exchanged a glance with Hyacinth.  "I'm enjoying this," he
seemed to say.
</P>

<P>
"Well," she announced, "I must be going in, too.  There'll be much to
see about."
</P>

<P>
Coronel was left alone with the most desperate lover in Araby.
</P>

<P>
"And now," said the Prince, "tell me what you are doing here."
</P>

<P>
Coronel put his arm in Udo's and walked him up and down the flagged
path.
</P>

<P>
"Your approaching marriage," he said, "is the talk of Araby.
Naturally I had to come here to see for myself what she was like.  My
dear Udo, she's charming; I congratulate you."
</P>

<P>
"Don't be a fool, Coronel.  I haven't the slightest intention of
marrying her."
</P>

<P>
"Then why have you told everybody that you are going to?"
</P>

<P>
"You know quite well I haven't told anybody.  There hasn't been a
single word about it mentioned until you pushed your way in just now."
</P>

<P>
"Ah, well, perhaps you hadn't heard about it.  But the Princess knows,
the Countess knows, and I know&mdash;yes, I think you may take our word for
it that it's true."
</P>

<P>
"I haven't the slightest intention&mdash;what do you keep clinging to my
arm like this for?
</P>

<P>
"My dear Udo, I'm so delighted to see you again.  Don't turn your back
on old friendships just because you have found a nobler and a
truer&mdash;&mdash;  Oh, very well, if you're going to drop all your former
friends, go on then.  But when <i>I'm</i> married, there will always be a
place for&mdash;&mdash;"
</P>

<P>
"Understand once and for all," said Udo angrily, "that I am <i>not</i>
getting married.  No, don't take my arm&mdash;we can talk quite well like
this."
</P>

<P>
"I am sorry, Udo," said Coronel meekly; "we seem to have made a
mistake.  But you must admit we found you in a very compromising
position."
</P>

<P>
"It wasn't in the least compromising," protested Udo indignantly.  "As
a matter of fact I was just telling her about that dragon I killed in
Araby last year."
</P>

<P>
"Ah, and who would listen to a hopeless story like that, but the woman
one was going to marry?"
</P>

<P>
"Once more, I am not going to marry her."
</P>

<P>
"Well, you must please yourself, but you have compromised her severely
with that story.  Poor innocent girl.  Well, let's forget about it.
And now tell me, how do you like Euralia?"
</P>

<P>
"I am returning to Araby this afternoon," said Udo stiffly.
</P>

<P>
"Well, perhaps you're right.  I hope that nothing will happen to you
on the way."
</P>

<P>
Udo, who was about to enter the Palace, turned round with a startled
look.
</P>

<P>
"What do you mean?"
</P>

<P>
"Well, something happened on the way here.  By the by, how did that
happen?  You never told me."
</P>

<P>
"Your precious Countess, whom you expect me to marry."
</P>

<P>
"How very unkind of her.  A nasty person to annoy."  He was silent for
a moment, and then added thoughtfully, "I suppose it <i>is</i> rather
annoying to think you're going to marry somebody whom you love very
much, and then find you're not going to."
</P>

<P>
Udo evidently hadn't thought of this.  He tried to show that he was
not in the least frightened.
</P>

<P>
"She couldn't do anything.  It was only by a lucky chance she did it
last time."
</P>

<P>
"Yes, but of course the chance might come again.  You'd have the thing
hanging over you always.  She's clever, you know; and I should never
feel quite safe if she were my enemy. . . .  Lovely flowers, aren't
they?  What's the name of this one?"
</P>

<P>
Udo dropped undecidedly into a seat.  This wanted thinking out.  The
Countess&mdash;what was wrong with her, after all? And she evidently adored
him.  Of course that was not surprising; the question was, was it fair
to disappoint one who had, perhaps, some little grounds for&mdash;&mdash;?
After all, he had been no more gallant than was customary from a
Prince and a gentleman to a beautiful woman.  It was her own fault if
she had mistaken his intentions.  Of course he ought to have left
Euralia long ago.  But he had stayed on, and&mdash;well, decidedly she was
beautiful&mdash;perhaps he had paid rather too much attention to that.  And
he had certainly neglected the Princess a little.  After all, again,
why not marry the Countess?  It was absurd to suppose there was
anything in Coronel's nonsense, but one never knew.  Not that he was
marrying her out of fear.  No; certainly not.  It was simply a
chivalrous whim on his part.  The poor woman had misunderstood him,
and she should not be disappointed.
</P>

<P>
"She seems fond of flowers," said Coronel.  "You ought to make the
Palace garden look beautiful between you."
</P>

<P>
"Now, understand clearly, Coronel, I'm not in the least frightened by
the Countess."
</P>

<P>
"My dear Udo, what a speech for a lover!  Of course you're not.  After
all, what you bore with such patience and dignity once, you can bear
again."
</P>

<P>
"That subject is distasteful to me.  I must ask you not to refer to
it.  If I marry the Countess&mdash;&mdash;"
</P>

<P>
"You'll be a very lucky man," put in Coronel.  "I happen to know that
the King of Euralia&mdash;however, she's chosen you, it seems.  Personally,
I can't make out what she sees in you.  What is it?"
</P>

<P>
"I should have thought it was quite obvious," said Udo with dignity.
"Well, Coronel, I think perhaps you are right and that it's my duty to
marry her."
</P>

<P>
Coronel shook him solemnly by the hand.
</P>

<P>
"I congratulate your Royal Highness.  I will announce your decision to
the Princess.  She will be much amu&mdash;much delighted."  And he turned
into the Palace.
</P>

<P>
Pity him, you lovers.  He had not seen Hyacinth for nearly ten
minutes.
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap20"></A>
<img src="images/0365X.jpg" alt="[Illustration: Detail of dark-haired girl in a pastoral scene]">
</p>

<H3 ALIGN="center">
CHAPTER XX
</H3>

<H3 ALIGN="center">
CORONEL KNOWS A GOOD STORY WHEN HE HEARS IT
</H3>

<P>
I quote (with slight alterations) from an epic by Charlotte Patacake,
a contemporary poet of the country:
</P>

<P class="poem">
    <i>King Merriwig the First rode back from war,</i> <BR>
    <i>As many other Kings had done before;</i><BR>
    <i>Five hundred men behind him were in sight</i><BR>
    <i>(Left-right, left-right, left-right, left-right, left-right).</i> <BR>
</P>

<P>
So far as is known, this was her only work, but she built up some
reputation on it, and Belvane, who was a good judge, had a high
opinion of her genius.
</P>

<P>
To be exact, there were only four hundred and ninety-nine men.  Henry
Smallnose, a bowman of considerable promise, had been left behind in
the enemy's country, the one casualty of war.  While spying out the
land in the early days of the invasion, he had been discovered by the
Chief Armourer of Barodia at full length on the wet grass searching
for tracks.  The Chief Armourer, a kindly man, had invited him to his
cottage, dried him and given him a warming drink, and had told him
that, if ever his spying took him that way again, he was not to stand
on ceremony, but come in and pay him a visit.  Henry, having caught a
glimpse of the Chief Armourer's daughter, had accepted without any
false pride, and had frequently dropped in to supper thereafter.  Now
that the war was over, he found that he could not tear himself away.
With King Merriwig's permission he was settling in Barodia, and with
the Chief Armourer's permission he was starting on his new life as a
married man.
</P>

<P>
As the towers of the castle came in sight, Merriwig drew a deep breath
of happiness.  Home again!  The hardships of the war were over; the
spoils of victory (wrapped up in tissue paper) were in his pocket;
days of honoured leisure were waiting for him.  He gazed at each
remembered landmark of his own beloved country, his heart overflowing
with thankfulness.  Never again would he leave Euralia!
</P>

<P>
How good to see Hyacinth again!  Poor little Hyacinth left all alone;
but there! she had had the Countess Belvane, a woman of great
experience, to help her.  Belvane!  Should he risk it?  How much had
she thought of him while he was away?  Hyacinth would be growing up
and getting married soon.  Life would be lonely in Euralia then,
unless&mdash;&mdash;  Should he risk it?
</P>

<P>
What would Hyacinth say?
</P>

<P>
She was waiting for him at the gates of the castle.  She had wanted
Coronel to wait with her, but he had refused.
</P>

<P class="noindent" align="center">
<a name="img0368"></a>
<img src="images/0368.jpg"
alt="[Illustration: As the towers of the Castle came in sight, Merriwig drew a deep breath of happiness, verso]">
<img src="images/0369.jpg"
alt="[Illustration: As the towers of the Castle came in sight, Merriwig drew a deep breath of happiness, recto]">
</P>

<P>
"We must offer the good news to him gradually," he said. "When a man
has just come back from a successful campaign, he doesn't want to find
a surprise like this waiting for him.  Just think&mdash;we don't even know
why the war is over&mdash;he must be longing to tell you that.  Oh, he'll
have a hundred things to tell you first; but then, when he says 'And
what's been happening here while I've been away? Nothing much, I
suppose?' then you can say&mdash;&mdash;"
</P>

<P>
"Then I shall say, 'Nothing much; only Coronel.'  And such a clever!"
</P>

<P>
"Oh, I have my ideas," said Coronel.  "Well, I'll be out of the way
somewhere.  I think I'll go for a walk in the forest.  Or shall I stay
here, in the Countess's garden, and amuse myself with Udo?  Anyhow,
I'll give you an hour alone together first."
</P>

<P>
The cavalcade drew up in front of the castle. Handkerchiefs fluttered
to them from the walls; trumpets were blown; hounds bayed.  Down the
steps came Hyacinth, all blue and gold, and flung herself into her
father's arms.
</P>

<P>
"My dear child," said Merriwig as he patted her soothingly.  "There,
there!  It's your old father come back again.  H'r'm.  There, there!"
He patted her again, as though it were she and not himself who was in
danger of breaking down.  "My little Hyacinth!  My own little girl!"
</P>

<P>
"Oh, Father, I <i>am</i> glad to have you back."
</P>

<P>
"There, there, my child.  Now I must just say a few words to my men,
and then we can tell each other all that has been happening."
</P>

<P>
He took a step forward and addressed his troops.
</P>

<P>
"Men of Euralia (<i>cheers</i>).  We have returned from a long and arduous
conflict (<i>cheers</i>) to the embraces (<i>loud cheers</i>) of our mothers and
wives and daughters (<i>prolonged cheering</i>)&mdash;as the case may be (<i>hear,
hear</i>).  In honour of our great victory I decree that, from now
onwards, to-morrow shall be observed as a holiday throughout Euralia
(<i>terrific cheering</i>).  I bid you all now return to your homes, and I
hope that you will find as warm a welcome there as I have found in
mine."  Here he turned and embraced his daughter again; and if his eye
travelled over her shoulder in the direction of Belvane's garden, it
is a small matter, and one for which the architect of the castle, no
doubt, was principally to blame.
</P>

<P>
There was another storm of cheers, the battle-cry of Euralia, "<i>Ho,
ho, Merriwig!</i>" was shouted from five hundred throats, and the men
dispersed happily to their homes.  Hyacinth and Merriwig went into the
Palace.
</P>

<P>
"Now, Father," said Hyacinth later on, when Merriwig had changed his
clothes and refreshed himself, "you've got to tell me all about it.  I
can hardly believe it's really over."
</P>

<P>
"Yes, yes.  It's all over," said Merriwig heartily.  "We shan't have
any trouble in <i>that</i> direction again, I fancy."
</P>

<P>
"Do tell me, did the King of Barodia apologise?"
</P>

<P>
"He did better than that, he abdicated."
</P>

<P>
"Why?"
</P>

<P>
"Well," said Merriwig, remembering just in time, "I&mdash;er&mdash;killed him."
</P>

<P>
"Oh, Father, how rough of you."
</P>

<P>
"I don't think it hurt him very much, my dear.  It was more a shock to
his feelings than anything else.  See, I have brought these home for
you."
</P>

<P>
He produced from his pocket a small packet in tissue paper.
</P>

<P>
"Oh, how exciting!  Whatever can it be?"
</P>

<P>
Merriwig unwrapped the paper, and disclosed a couple of ginger
whiskers, neatly tied up with blue ribbon.
</P>

<P>
"Father!"
</P>

<P>
He picked out the left one, <i>fons et origo</i> (if he had known any
Latin) of the war, and held it up for Hyacinth's inspection.
</P>

<P>
"There, you can see the place where Henry Smallnose's arrow bent it.
By the way," he added, "Henry is marrying and settling down in
Barodia.  It is curious," he went on, "how after a war one's thoughts
turn to matrimony."  He glanced at his daughter to see how she would
take this, but she was still engrossed with the whiskers.
</P>

<P>
"What am I going to do with them, Father?  I can't plant them in the
garden."
</P>

<P>
"I thought we might run them up the flagstaff, as we did in Barodia."
</P>

<P>
"Isn't that a little unkind now that the poor man's dead?"
</P>

<P>
Merriwig looked round him to see that there were no eavesdroppers.
</P>

<P>
"Can you keep a secret?" he asked mysteriously.
</P>

<P>
"Of course," said Hyacinth, deciding at once that it would not matter
if she only told Coronel.
</P>

<P>
"Well, then, listen."
</P>

<P>
He told her of his secret journey to the King of Barodia's tent; he
told her of the King of Barodia's letter; he told her more fully of
his early duel with the King; he told her everything that he had said
and done; and everything that everybody else had said and done to him;
and his boyish pleasure in it all was so evident and so innocent, that
even a stranger would have had nothing more reproachful for him than a
smile.  To Hyacinth he seemed the dearest of fathers and the most
wonderful of kings.
</P>

<P>
And by and by the moment came of which Coronel had spoken.
</P>

<P>
"And now," said Merriwig, "tell me what you have all been doing with
yourselves here.  Nothing much, I suppose?"
</P>

<P>
He waited nervously, wondering if Hyacinth would realise that "all"
was meant to include more particularly Belvane.
</P>

<P>
Hyacinth drew a stool up to her father's chair and sat down very close
to him.
</P>

<P>
"Father," she said, stroking his hand where it rested on his knee, "I
<i>have</i> got some news for you."
</P>

<P>
"Nothing about the Coun&mdash;nothing serious, I hope," said Merriwig, in
alarm.
</P>

<P>
"It's rather serious, but it's rather nice.  Father, dear, would you
mind <i>very</i> much if I got married soon?"
</P>

<P>
"My dear, you shall get married as soon as you like.  Let me see,
there were six or seven Princes who came about it only the other day.
I sent them off on adventures of some kind, but&mdash;dear me, yes, they
ought to have been back by now.  I suppose you haven't heard anything
of them?"
</P>

<P>
"No, Father," said Hyacinth, with a little smile.
</P>

<P>
"Ah, well, no doubt they were unsuccessful.  No matter, dear, we can
easily find you plenty more suitors.  Indeed, the subject has been
very near my thoughts lately.  We'll arrange a little competition, and
let them know in the neighbouring countries; there'll be no lack of
candidates.  Let me see, there's that seven-headed bull; he's getting
a little old now, but he was good enough for the last one.  We
might&mdash;&mdash;"
</P>

<P>
"I don't want a suitor," said Hyacinth softly.  "I have one."
</P>

<P>
Merriwig leant forward with eagerness.
</P>

<P>
"My dear, this is indeed news.  Tell me all about it.  Upon what quest
did you send him?"
</P>

<P>
Hyacinth had felt this coming.  Had she lived in modern times she
would have expected the question, "What is his income?"  A man must
prove his worth in some way.
</P>

<P>
"I haven't sent him away at all yet," she said; "he's only just come.
He's been very kind to me, and I'm sure you'll love him."
</P>

<P>
"Well, well, we'll arrange something for him.  Perhaps that bull I was
speaking of&mdash;&mdash;  By the way, who is he?"
</P>

<P>
"He comes from Araby, and his name is&mdash;&mdash;"
</P>

<P>
"Udo, of course.  Why didn't I think of him?  An excellent
arrangement, my dear."
</P>

<P>
"It isn't Udo, I'm afraid, Father.  It's Coronel."
</P>

<P>
"And who might Coronel be?" said the King, rather sternly.
</P>

<P>
"He's&mdash;he's&mdash;well, he's&mdash;&mdash;  Here he is, Father."  She ran up to him
impulsively as he came in at the door.  "Oh, Coronel, you're just in
time; do tell Father who you are."
</P>

<P>
Coronel bowed profoundly to the King.
</P>

<P>
"Before I explain myself, your Majesty," he said, "may I congratulate
your Majesty on your wonderful victory over the Barodians?  From the
little I have gathered outside, it is the most remarkable victory that
has ever occurred.  But of course I am longing to hear the full story
from your Majesty's own lips.  Is it a fact that your Majesty made his
way at dead of night to the King of Barodia's own tent and challenged
him to mortal combat and slew him?"  There was an eagerness, very
winning, in his eyes as he asked it; he seemed to be envying the King
such an adventure&mdash;an adventure after his own heart.
</P>

<P>
Merriwig was in an awkward position.  He wondered for a moment whether
to order his daughter out of the room. "Leave us, my child," he would
say.  "These are matters for men to discuss."  But Hyacinth would know
quite well why she had been sent out, and would certainly tell Coronel
the truth of the matter afterwards.
</P>

<P>
It really looked as if Coronel would have to be let into the secret
too.  He cleared his throat noisily by way of preparation.
</P>

<P>
"There are certain state reasons," he said with dignity, "why that
story has been allowed to get about."
</P>

<P>
"Pardon, your Majesty.  I have no wish to&mdash;&mdash;"
</P>

<P>
"But as you know so much, you may as well know all.  It happened like
this."  Once more he told the story of his midnight visit, and of the
King's letter to him.
</P>

<P>
"But, your Majesty," cried Coronel, "it is more wonderful than the
other.  Never was such genius of invention, such brilliance and daring
of execution."
</P>

<P>
"So you like it," said Merriwig, trying to look modest.
</P>

<P>
"I love it."
</P>

<P>
"I knew he'd love it," put in Hyacinth.  "It's just the sort of story
that Coronel would love.  Tell him about how you fought the King at
the beginning of the war, and how you pretended to be a swineherd, and
how&mdash;"
</P>

<P>
Could any father have resisted?  In a little while Hyacinth and
Coronel were seated eagerly at his feet, and he was telling once more
the great story of his adventures.
</P>

<P>
"Well, well," said the King at the end of it, when he had received
their tribute of admiration.  "Those are just a few of the little
adventures that happen in war time."  He turned to Coronel.  "And so
you, I understand, wish to marry my daughter?"
</P>

<P>
"Does that surprise your Majesty?"
</P>

<P>
"Well, no, it doesn't.  And she, I understand, wishes to marry you."
</P>

<P>
"Yes, please, Father."
</P>

<P>
"That," said Coronel simply, "is much more surprising."
</P>

<P>
Merriwig, however, was not so sure of that.  He liked the look of
Coronel, he liked his manner, and he saw at once that he knew a good
story&mdash;when he heard one.
</P>

<P>
"Of course," he said, "you'll have to win her."
</P>

<P>
"Anything your Majesty sets me to do.  It's as well," he added with a
disarming smile, "that you cannot ask for the whiskers of the King of
Barodia.  There is only one man who could have got those."
</P>

<P>
Truly an excellent young man.
</P>

<P>
"Well, we'll arrange something," said Merriwig, looking pleased.
"Perhaps your Prince Udo would care to be a competitor too."
</P>

<P>
Hyacinth and Coronel interchanged a smile.
</P>

<P>
"Alas, Father," she said, "his Royal Highness is not attracted by my
poor charms."
</P>

<P>
"Wait till he has seen them, my dear," said Merriwig with a chuckle.
</P>

<P>
"He has seen them, Father."
</P>

<P>
"What?  You invited him here?  Tell me about this, Hyacinth.  He came
to stay with you and he never&mdash;&mdash;"
</P>

<P>
"His Royal Highness," put in Coronel, "has given his affections to
another."
</P>

<P>
"Aha!  So that's the secret.  Now I wonder if I can guess who she is.
What do you say to the Princess Elvira of Tregong?  I know his father
had hopes in that direction."
</P>

<P>
Hyacinth looked round at Coronel as if appealing for his support.  He
took a step towards her.
</P>

<P>
"No, it's not the Princess Elvira," said Hyacinth, a little nervously.
</P>

<P>
The King laughed good-humouredly.
</P>

<P>
"Ah, well, you must tell me," he said.
</P>

<P>
Hyacinth put out her hand, and Coronel pressed it encouragingly.
</P>

<P>
"His Royal Highness Prince Udo," she said, "is marrying the Countess
Belvane."
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap21"></A>
<img src="images/0385X.jpg" alt="[Illustration: A man surrounded by clouds of smoke]">
</p>

<H3 ALIGN="center">
CHAPTER XXI
</H3>

<H3 ALIGN="center">
A SERPENT COMING AFTER UDO
</H3>

<P>
Belvane had now had twenty-four hours in which to think it over.
</P>

<P>
Whatever her faults, she had a sense of humour.  She could not help
smiling to herself as she thought of that scene in the garden.
However much she regretted her too hasty engagement, she was sure Udo
regretted it still more.  If she gave him the least opportunity he
would draw back from it.
</P>

<P>
Then why not give him the opportunity?  "My dear Prince Udo, I'm
afraid I mistook the nature of my feelings"&mdash;said, of course, with
downcast head and a maidenly blush.  Exit Udo with haste, enter King
Merriwig.  It would be so easy.
</P>

<P>
Ah, but then Hyacinth would have won.  Hyacinth had forced the
engagement upon her; even if it only lasted for twenty-four hours, so
long as it was a forced engagement, Hyacinth would have had the better
of her for that time. But if she welcomed the engagement, if she
managed in some way to turn it to account, to make it appear as if she
had wanted it all the time, then Hyacinth's victory would be no
victory at all, but a defeat.
</P>

<P>
Marry Udo, then, as if willingly?  Yes, but that was too high a price
to pay.  She was by this time thoroughly weary of him and besides, she
had every intention of marrying the King of Euralia.  To pretend to
marry him until she brought the King in open conflict with him, and
then having led the King to her feet to dismiss the rival who had
served her turn&mdash;that was her only wise course.
</P>

<P>
She did not come to this conclusion without much thought. She composed
an Ode to Despair, an Elegy to an Unhappy Woman, and a Triolet to
Interfering Dukes, before her mind was made up.  She also considered
very seriously what she would look like in a little cottage in the
middle of the forest, dressed in a melancholy grey and holding
communion only with the birds and trees; a life of retirement away
from the vain world; a life into which no man came.  It had its
attractions, but she decided that grey did not suit her.
</P>

<P>
She went down to her garden and sent for Prince Udo.  At about the
moment when the King was having the terrible news broken to him, Udo
was protesting over the sundial that he loved Belvane and Belvane
only, and that he was looking forward eagerly to the day when she
would make him the happiest of men.  So afraid was he of what might
happen to him on the way back to Araby.
</P>

<P>
"The Countess Belvane!" cried Merriwig.  "Prince Udo marry the
Countess Belvane!  I never heard such a thing in my life."  He glared
at them one after the other as if it were their fault&mdash;as indeed it
was.  "Why didn't you tell me this before, Hyacinth?"
</P>

<P>
"It was only just announced, Father."
</P>

<P>
"Who announced it?"
</P>

<P>
"Well&mdash;er&mdash;Udo did," said Coronel.
</P>

<P>
"I never heard of anything so ridiculous in my life!  I won't have
it!"
</P>

<P>
"But, Father, don't you think she'd make a very good Queen?"
</P>

<P>
"She'd make a wonderful&mdash;that has nothing to do with it.  What I feel
so strongly about is this.  For month after month I am fighting in a
strange country.  After extraordinary scenes of violence and&mdash;peril&mdash;I
come back to my own home to enjoy the&mdash;er&mdash;fruits of victory.  No
sooner do I get inside my door than I have all this thrust upon me."
</P>

<P>
"All what, Father?" said Hyacinth innocently.
</P>

<P>
"All <i>this</i>," said the King, with a circular movement of his hand.
"It's too bad; upon my word it is.  I won't have it.  Now mind,
Hyacinth, I <i>won't</i> have it.
</P>

<P>
"But, Father, how can I help it?"
</P>

<P>
Merriwig paid no attention to her.
</P>

<P>
"I come home," he went on indignantly, "fresh from the&mdash;er&mdash;spoils of
victory to what I thought was my own peaceful&mdash;er&mdash;home.  And what do
I find?  Somebody here wants to marry somebody there, and somebody
else over there wants to marry somebody else over here; it's
impossible to mention any person's name, in even the most casual way,
without being told they are going to get married, or some nonsense of
that sort.  I'm very much upset about it."
</P>

<P>
"Oh, Father!" said Hyacinth penitently.  "Won't you see the Countess
yourself and talk to her?"
</P>

<P>
"To think that for weeks I have been looking forward to my return home
and that now I should be met with this!  It has quite spoilt my day."
</P>

<P>
"Father!" cried Hyacinth, coming towards him with outstretched hands.
</P>

<P>
"Let me send for her ladyship," began Coronel; "perhaps she&mdash;&mdash;"
</P>

<P>
"No, no," said Merriwig, waving them away.  "I am very much displeased
with you both.  What I have to do, I can do quite well by myself."
</P>

<P>
He strode out and slammed the door behind him.
</P>

<P>
Hyacinth and Coronel looked at each other blankly.
</P>

<P>
"My dear," said Coronel, "you never told me he was as fond of her as
that."
</P>

<P>
"But I had no idea!  Coronel, what can we do now about it?  Oh, I want
him to marry her now.  He's quite right&mdash;she'll make a wonderful
Queen.  Oh, my dear, I feel I want everybody to be as happy as we're
going to be."
</P>

<P>
"They can't be that, but we'll do our best for them.  I can manage Udo
all right.  I only have to say 'rabbits' to him, and he'll do anything
for me.  Hyacinth, I don't believe I've ever kissed you in this room
yet, have I?  Let's begin now."
</P>

<P>
Merriwig came upon the other pair of lovers in Belvane's garden.  They
were sharing a seat there, and Udo was assuring the Countess that he
was her own little Udo-Wudo, and that they must never be away from
each other again. The King put his hand in front of his eyes for a
moment as if he could hardly bear it.
</P>

<P>
"Why, it's his Majesty," said Belvane, jumping up.  She gave him a
deep curtsey and threw in a bewitching smile on the top of it;
formality or friendliness, he could take his choice.  "Prince Udo of
Araby, your Majesty."  She looked shyly at him and added, "Perhaps you
have heard."
</P>

<P>
"I have," said the King gloomingly.  "How do you do," he added in a
melancholy voice.
</P>

<P>
Udo declared that he was in excellent health at present, and would
have gone into particulars about it had not the King interrupted.
</P>

<P>
"Well, Countess," he said, "this is strange news to come back to.
Shall I disturb you if I sit down with you for little?"
</P>

<P>
"Oh, your Majesty, you would honour us.  Udo, dear, have you seen the
heronry lately?"
</P>

<P>
"Yes," said Udo.
</P>

<P>
"It looks so sweet just about this time of the afternoon."
</P>

<P>
"It does," said Udo.
</P>

<P>
Belvane gave a little shrug and turned to the King.
</P>

<P>
"I'm so longing to hear all your adventures," she murmured
confidingly.  "I got all your messages; it was so good of you to
remember me."
</P>

<P>
"Ah," said Merriwig reproachfully, "and what do I find when I come
back?  I find&mdash;&mdash;"  He broke off, and indicated in pantomime with his
eyebrows that he could explain better what he had found if Udo were
absent.
</P>

<P>
"Udo, dear," said Belvane, turning to him, "have you seen the kennels
lately?"
</P>

<P>
"Yes," said Udo.
</P>

<P>
"They look rather sweet just about this time," said Merriwig.
</P>

<P>
"Don't they?" said Udo.
</P>

<P>
"But I am so longing to hear," said Belvane, "how your Majesty
defeated the King of Barodia.  Was it your Majesty's wonderful spell
which overcame the enemy?"
</P>

<P>
"You remember that?"
</P>

<P>
"Remember it?  Oh, your Majesty!  '<i>Bo boll&mdash;&mdash;</i>'  Udo, dear, wouldn't
you like to see the armoury?"
</P>

<P>
"No," said Udo.
</P>

<P>
"There are a lot of new things in it that I brought back from
Barodia," said Merriwig hopefully.
</P>

<P>
"A lot of new things," explained Belvane.
</P>

<P>
"I'll see them later on," said Udo.  "I dare say they'd look better in
the evening."
</P>

<P>
"Then you shall show <i>me</i>, your Majesty," said Belvane. "Udo, dear,
you can wait for me here."
</P>

<P>
The two of them moved off down the path together (Udo taken by
surprise), and as soon as they were out of sight, tiptoed across the
lawn to another garden seat, Belvane leading the way with her finger
to her lips, and Merriwig following with an exaggerated caution which
even Henry Smallnose would have thought overdone.
</P>

<P>
"He is a little slow, isn't he, that young man?" said the King, as
they sat down together.  "I mean he didn't seem to understand&mdash;"
</P>

<P>
"He's such a devoted lover, your Majesty.  He can't bear to be out of
my sight for a moment."
</P>

<P>
"Oh, Belvane, this is a sad homecoming.  For month after month I have
been fighting and toiling, and planning and plotting and then&mdash;&mdash;  Oh,
Belvane, we were all so happy together before the war."
</P>

<P>
Belvane remembered that once she and the Princess and Wiggs had been
so happy together, and that Udo's arrival had threatened to upset it
all.  One way and another, Udo had been a disturbing element in
Euralia.  But it would not do to let him go just yet.
</P>

<P>
"Aren't we still happy together?" she asked innocently.  "There's her
Royal Highness with her young Duke, and I have my dear Udo, and your
Majesty has the&mdash;the Lord Chancellor&mdash;and all your Majesty's faithful
subjects."
</P>

<P>
His Majesty gave a deep sigh.
</P>

<P class="noindent" align="center">
<a name="img0396"></a>
<img src="images/0396.jpg" alt="[Illustration: Belvane leading the way with her finger to her lips]">
<a name="img0397"></a>
<img src="images/0397.jpg" alt="[Illustration: Merriwig following with an exaggerated caution]">
</P>

<P>
"I am a very lonely man, Belvane.  When Hyacinth leaves me I shall
have nobody left."
</P>

<P>
Belvane decided to risk it.
</P>

<P>
"Your Majesty should marry again," she said gently.
</P>

<P>
He looked unutterable things at her.  He opened his mouth with the
intention of doing his best to utter some of them, when&mdash;&mdash;
</P>

<P>
"Not before Udo," said Belvane softly.
</P>

<P>
Merriwig got up indignantly and scowled at the Prince as the latter
hurried over the lawn towards them.
</P>

<P>
"Well, really," said Merriwig, "I never knew such a place.  One simply
can't&mdash;&mdash;  Ah, your Royal Highness, have you seen our armoury?  I
should say," he corrected himself as he caught Belvane's reproachful
look, "have <i>we</i> seen our armoury?  We have.  Her ladyship was much
interested."
</P>

<P>
"I have no doubt, your Majesty."  He turned to Belvane.  "You will be
interested in our armoury at home, dear."
</P>

<P>
She gave a quick glance at the King to see that he was looking, and
then patted Udo's hand tenderly.
</P>

<P>
"Home," she said lovingly, "how sweet it sounds!"
</P>

<P>
The King shivered as if in pain, and strode quickly from them.
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
"Your Majesty sent for me," said Coronel.
</P>

<P>
The King stopped his pacings and looked round as Coronel came into the
library.
</P>

<P>
"Ah, yes, yes," he said quickly.  "Now sit down there and make
yourself comfortable.  I want to talk to you about this marriage."
</P>

<P>
"Which one, your Majesty?"
</P>

<P>
"Which one?  Why, of course, yours&mdash;that is to say,
Belvane's&mdash;or&mdash;rather&mdash;&mdash;"  He came to a stop in front of Coronel and
looked at him earnestly.  "Well, in a way, both."
</P>

<P>
Coronel nodded.
</P>

<P>
"You want to marry my daughter," Merriwig went on.  "Now it is
customary, as you know, that to the person to whom I give my daughter,
I give also half my kingdom.  Naturally before I make this sacrifice I
wish to be sure that the man to whom&mdash;well, of course, you
understand."
</P>

<P>
"That he is worthy of the Princess Hyacinth," said Coronel.  "Of
course he couldn't be," he added with a smile.
</P>

<P>
"<i>And</i> worthy of half the kingdom," amended Merriwig. "That he should
prove himself this is also, I think, customary."
</P>

<P>
"Anything that your Majesty suggests&mdash;&mdash;"
</P>

<P>
"I am sure of it."
</P>

<P>
He drew up a chair next to Coronel's, and sitting down in it, placed
his hand upon his knees and explained the nature of the trial which
was awaiting the successful suitor.
</P>

<P>
"In the ordinary way," he began, "I should arrange something for you
with a dragon or what-not in it.  The knowledge that some such ordeal
lies before him often enables a suitor to discover, before it is too
late, that what he thought was true love is not really the genuine
emotion.  In your case I feel that an ordeal of this sort is not
necessary."
</P>

<P>
Coronel inclined his head gracefully.
</P>

<P>
"I do not doubt your valour, and from you therefore I ask a proof of
your cunning.  In these days cunning is perhaps the quality of all
others demanded of a ruler.  We had an excellent example of that," he
went on carelessly, "in the war with Barodia that is just over, where
the whole conflict was settled by a little idea which&mdash;&mdash;"
</P>

<P>
"A very wonderful idea, your Majesty."
</P>

<P>
"Well, well," said Merriwig, looking very pleased.  "It just happened
to come off, that's all.  But that is what I mean when I say that
cunning may be of even more importance than valour.  In order to win
the hand of my daughter and half my kingdom, it will be necessary for
you to show a cunning almost more than human."
</P>

<P>
He paused, and Coronel did his best in the interval to summon up a
look of superhuman guile into his very frank and pleasant countenance.
</P>

<P>
"You will prove yourself worthy of what you ask me for," said Merriwig
solemnly, "by persuading Prince Udo to return to Araby&mdash;alone."
</P>

<P>
Coronel gasped.  The thing was so easy that it seemed almost a shame
to accept it as the condition of his marriage.  To persuade Udo to do
what he was only longing to do, did not call for any superhuman
qualities of any kind.  For a moment he had an impulse to tell the
King so, but he suppressed it.  "After all," he thought, "if the King
wants cunning, and if I make a great business of doing something
absurdly easy, then he is getting it."
</P>

<P>
Merriwig, simple man, mistook his emotions.
</P>

<P>
"I see," he said, "that you are appalled by the difficulty of the
ordeal in front of you.  You may well be so.  You have known his Royal
Highness longer than I have, but even in our short acquaintance I have
discovered that he takes a hint with extraordinary slowness.  To bring
it home to him with the right mixture of tact and insistence that
Araby needs his immediate presence&mdash;alone&mdash;may well tax the most
serpentine of minds."
</P>

<P>
"I can but try it," said the serpentine one simply.
</P>

<P>
The King jumped up and shook him warmly by the hand.
</P>

<P>
"You think you can do it?" he said excitedly.
</P>

<P>
"If Prince Udo does not start back to Araby to-morrow&mdash;&mdash;"
</P>

<P>
"Alone," said Merriwig.
</P>

<P>
"Alone&mdash;then I shall have failed in my task."
</P>

<P>
     &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;* * * * *<BR>
</P>

<P>
"My dear," said the King to his daughter as she kissed him good-night
that evening, "I believe you are going to marry a very wise young
man."
</P>

<P>
"Of course I am, Father."
</P>

<P>
"I only hope you'll be as happy with him as I shall be with&mdash;as I was
with your mother.  Though how he's going to bring it off," he added to
himself, "is more than I can think."
</P>

<BR><BR><BR>

<p class="noindent" align="center">
<A NAME="chap22"></A>
<img src="images/0405X.jpg" alt="[Illustration: Same image as for chapter 20]">
</p>

<H3 ALIGN="center">
CHAPTER XXII
</H3>

<H3 ALIGN="center">
THE SEVENTEEN VOLUMES GO BACK AGAIN
</H3>

<P>
King Merriwig of Eastern Euralia sat at breakfast on his castle walls.
He lifted the gold cover from the gold dish in front of him, selected
a trout, and conveyed it carefully to his gold plate.  When you have
an aunt&mdash;&mdash;  But I need not say that again.
</P>

<P>
King Coronel of Western Euralia sat at breakfast on <i>his</i> castle
walls.  He lifted the gold cover from the gold dish in front of him,
selected a trout, and conveyed it carefully to his gold plate.  When
your wife's father has an aunt&mdash;&mdash;
</P>

<P>
Prince Udo of Araby sat at breakfast&mdash;&mdash;  But one must draw the line
somewhere.  I refuse to follow Udo through any more meals.  Indeed, I
think there has been quite enough eating and drinking in this book
already.  Quite enough of everything in fact; but the time has nearly
come to say good-bye.
</P>

<P>
Let us speed the Prince of Araby first.  His departure from Euralia
was sudden; five minutes' conversation with Coronel convinced him that
there had been a mistake about Belvane's feelings for him, and that he
could leave for Araby in perfect safety.
</P>

<P>
"You must come and see us again," said Merriwig heartily, as he shook
him by the hand.
</P>

<P>
"Yes, do," said Hyacinth.
</P>

<P>
There are two ways of saying this sort of thing, and theirs was the
second way.  So was Udo's, when he answered that he would be
delighted.
</P>

<P>
It was just a week later that the famous double wedding was celebrated
in Euralia.  As an occasion for speech-making by King Merriwig and
largesse-throwing by Queen Belvane it demanded and (got) a whole
chapter to itself in Roger's History.  I have Roger on my side at
last.  The virtues he denied to the Countess he cannot but allow to
the Queen.
</P>

<P>
Nor could Hyacinth resist her any longer.  Belvane upon her palfrey,
laughter in her eyes and roses in her cheeks, her lips slightly parted
with eagerness as she flings her silver to the crowd, adorably
conscious of her childishness and yet glorifying in it, could have no
enemies that day.
</P>

<P>
"She is a dear," said Hyacinth to Coronel.  "She will make a wonderful
Queen."
</P>

<P>
"I know a Queen worth two of her," said Coronel.
</P>

<P>
"But you do admire her, don't you?"
</P>

<P>
"Not particularly."
</P>

<P>
"Oh, Coronel, you must," said Hyacinth, but she felt very happy all
the same.
</P>

<P>
They rode off the next day to their kingdom.  The Chancellor had had
an exciting week; for seven successive evenings he had been extremely
mysterious and reserved to his wife, but now his business was finished
and King Merriwig reigned over Eastern Euralia and King Coronel over
the West.
</P>

<P>
Let us just take a look at Belvane's diary before we move on to the
last scene.
</P>

<P>
"<i>Thursday, September 15th</i>," it says.  "<i>Became good.</i>"
</P>

<P>
Now for the last scene.
</P>

<P>
King Merriwig sat in Belvane's garden.  They had spent the morning
revising their joint book of poetry for publication.  The first set of
verses was entirely Merriwig's own.  It went like this:
</P>

<P class="poem">
           <i>Bo, boll, bill, bole.</i> <BR>
           <i>Wo, woll, will, wole.</i> <BR>
</P>

<P>
A note by the authors called attention to the fact that it could be
begun from either end.  The rest of the poems were mainly by Belvane,
Merriwig's share in them consisting of a "Capital," or an "I like
that," when they were read out to him; but an epic commonly attributed
to Charlotte Patacake had crept in somehow.
</P>

<P>
"A person to see your Majesty," said a flunkey, appearing suddenly.
</P>

<P>
"What sort of person?" asked Merriwig.
</P>

<P>
"A sort of person, your Majesty."
</P>

<P>
"See him here, dear," said Belvane, as she got up.  "I have things to
do in the Palace."
</P>

<P>
She left him; and by and by the flunkey returned with the stranger.
He was a pleasant-looking person with a round clean-shaven face;
something in the agricultural way, to judge from his clothes.
</P>

<P>
"Well?" said Merriwig.
</P>

<P>
"I desire to be your Majesty's swineherd," said the other.
</P>

<P>
"What do you know of swineherding?"
</P>

<P>
"I have a sort of natural aptitude for it, your Majesty, although I
have never actually been one."
</P>

<P>
"My own case exactly.  Now then, let me see&mdash;how would you&mdash;&mdash;"
</P>

<P>
The stranger took out a large red handkerchief and wiped his forehead.
</P>

<P>
"You propose to ask me a few questions, your Majesty?"
</P>

<P>
"Well, naturally, I&mdash;&mdash;"
</P>

<P>
"Let me beg of you not to.  By all you hold sacred let me implore you
not to confuse me with questions."  He drew himself up and thumped his
chest with his fist.  "I have a feeling for swineherding; it is
enough."
</P>

<P>
Merriwig began to like the man; it was just how he felt about the
thing himself.
</P>

<P>
"I once carried on a long technical conversation with a swineherd," he
said reminiscently, "and we found we had much in common.  It is an
inspiring life."
</P>

<P>
"It was in just that way," said the stranger, "that I discovered my
own natural bent towards it."
</P>

<P>
"How very odd!  Do you know, there's something about your face that I
seem to recognise?"
</P>

<P>
The stranger decided to be frank.
</P>

<P>
"I owe this face to you," he said simply.
</P>

<P>
Merriwig looked startled.
</P>

<P>
"In short," said the other, "I am the late King of Barodia."
</P>

<P>
Merriwig gripped his hand.
</P>

<P class="noindent" align="center">
<a name="img0412"></a>
<img src="images/0412.jpg"
alt="[Illustration: He was a pleasant-looking person, with a round clean-shaven face, verso]">
<img src="images/0413.jpg"
alt="[Illustration: He was a pleasant-looking person, with a round clean-shaven face, recto]">
</P>

<P>
"My dear fellow," he said.  "My very dear fellow, of course you are.
Dear me, how it brings it all back.  And&mdash;may I say&mdash;what an
improvement.  Really, I'm delighted to see you.  You must tell me all
about it.  But first some refreshment."
</P>

<P>
At the word "refreshment" the late King of Barodia broke down
altogether, and it was only Merriwig's hummings and hawings and
thumpings on the back and (later on) the refreshment itself which kept
him from bursting into tears.
</P>

<P>
"My dear friend," he said, as he wiped his mouth for the last time,
"you have saved me."
</P>

<P>
"But what does it all mean?" asked Merriwig in bewilderment.
</P>

<P>
"Listen and I will tell you,"
</P>

<P>
He told himself of the great resolution to which he had come on that
famous morning when he awoke to find himself whiskerless.  Barodia had
no more use for him now as a King, and he on his side was eager to
carve out for himself a new life as a swineherd.
</P>

<P>
"I had a natural gift," he said plaintively, "an instinctive feeling
for it.  I know I had.  Whatever they said about it afterwards&mdash;and
they said many hard things&mdash;I was certain that I had that feeling.  I
had proved it, you know; there couldn't be any mistake."
</P>

<P>
"Well?"
</P>

<P>
"Ah, but they laughed at me.  They asked me confusing questions;
niggling little questions about the things swine ate and&mdash;and things
like that.  The great principles of swineherding, the&mdash;what I may call
the art of herding swine, the whole theory of shepherding pigs in a
broad-minded way, all this they ignored.  They laughed at me and
turned me out with jeers and blows&mdash;to starve."
</P>

<P>
Merriwig patted him sympathetically, and pressed some more food on
him.
</P>

<P>
"I ranged over the whole of Barodia.  Nobody would take me in.  It is
a terrible thing, my dear Merriwig, to begin to lose faith in
yourself.  I had to tell myself at last that perhaps there was
something about Barodian swine which made them different from those of
any other country.  As a last hope I came to Euralia; if here too I
was spurned, then I should know that&mdash;&mdash;"
</P>

<P>
"Just a moment," said Merriwig, breaking in eagerly.  "Who was this
swineherd that you talked to&mdash;&mdash;"
</P>

<P>
"I talked to so many," said the other sadly.  "They all scoffed at
me."
</P>

<P>
"No, but the first one; the one that showed you that you had a bent
towards it.  Didn't you say that&mdash;&mdash;"
</P>

<P>
"Oh, that one.  That was at the beginning of our war.  Do you remember
telling me that your swineherd had an invisible cloak?  It was he
that&mdash;&mdash;"
</P>

<P>
Merriwig looked at him sadly and shook his head.
</P>

<P>
"My poor friend," he said, "it was me."
</P>

<P>
They gazed at each other earnestly.  Each of them was going over in
his mind the exact details of that famous meeting.
</P>

<P>
"Yes," they murmured together, "it was us."
</P>

<P>
The King of Barodia's mind raced on through all the bitter months that
had followed; he shivered as he thought of the things he had said; the
things that had been said to him seemed of small account now.
</P>

<P>
"Not even a swineherd!" he remarked.
</P>

<P>
"Come, come," said Merriwig, "look on the bright side; you can always
be a King again."
</P>

<P>
The late King of Barodia shook his head.
</P>

<P>
"It's a come down to a man with any pride," he said.  "No, I'll stick
to my own job.  After all, I've been learning these last weeks; at any
rate I know that what I do know isn't worth knowing, and that's
something."
</P>

<P>
"Then stay with me," said Merriwig heartily.  "My swineherd will teach
you your work, and when he retires you can take it on."
</P>

<P>
"Do you mean it?"
</P>

<P>
"Of course I do.  I shall be glad to have you about the place.  In the
evening, when the pigs are asleep, you can come in and have a chat
with us."
</P>

<P>
"Bless you," said the new apprentice; "bless you, your Majesty."
</P>

<P>
They shook hands on it.
</P>

<P>
"My dear," said Merriwig to Belvane that evening, "you haven't married
a very clever fellow.  I discovered this afternoon that I'm not even
as clever as I thought I was."
</P>

<P>
"You don't want cleverness in a King," said Belvane, smiling lovingly
at him, "or in a husband."
</P>

<P>
"What do you want then?"
</P>

<P>
"Just dearness," said Belvane.
</P>

<P>
</P>

<P>
</P>

<P>
And now my story is done.  With a sigh I unload the seventeen volumes
of Euralian History from my desk, carrying them one by one across the
library and placing them carefully in the shelf which has been built
for them.  For some months they have stood a rampart between me and
the world, behind which I have lived in far-off days with Merriwig and
Hyacinth and my Lady Belvane.  The rampart is gone, and in the bright
light of to-day which streams on to my desk the vision slowly fades.
Once on a time . .
</P>

<P>
Yet I see one figure clearly still.  He is tall and thin, with a white
peaked face of which the long inquisitive nose is the outstanding
feature.  His hair is lank and uncared for; his russet smock, tied in
at the waist, wants brushing; his untidy cross-gartered hose shows up
the meagerness of his legs.  No knightly figure this, yet I look upon
him very tenderly.  For it is Roger Scurvilegs on his way to the
Palace for news.
</P>

<P>
To Roger too I must say good-bye.  I say it not without remorse, for I
feel that I have been hard upon the man to whom I owe so much.
Perhaps it will not be altogether good-bye; in his seventeen volumes
there are many other tales to be found.  Next time (if there be a next
time) I owe it to Roger to stand aside and let him tell the story more
in his own way.  I think he would like that.
</P>

<P>
But it shall not be a story about Belvane.  I saw Belvane (or some one
like her) at a country house in Shropshire last summer, and I know
that Roger can never do her justice.
</P>

<P class="noindent" align="center">
<a name="img0420X"></a>
<img src="images/0420X.jpg" alt="[Illustration: Roger Scurvilegs]">
</p>

<BR><BR><BR><BR>

<P class="noindent" align="center">
<img src="images/0422.jpg" alt="[Illustration: Back endpaper, verso]">
<img src="images/0423.jpg" alt="[Illustration: Back endpaper, recto]">
</P>

<BR><BR><BR>

<p class="noindent" align="right">
<img src="images/0424.jpg" alt="[Illustration: Back cover]">
</p>

<BR><BR><BR><BR>
`
