# Fuji-Llama 
This is a Llama Card game server written in GO. 
This is my application written in GO, 

so has been a learnign exersise in Go so probaly not the best coding example. 

But you can play Llama just install Go onto your computer and download the code and run it

You can  play Llama in your terminal program as 1 player (Me) against up to 5 AI bots

# The game of Llama 

In the L.L.A.M.A. card game, the goal is to get rid of all your cards by playing them in sequence, 
either matching the top card of the discard pile or playing one number higher, 
including llamas, which can be played on 6s or other llamas. 
Players can also draw cards or quit the round, 
with penalties assessed based on the cards left in hand at the end of the round. 
The game continues until a player reaches 40 points, with the lowest score winning. 

Here's a more detailed breakdown:
# Game Setup:
Players: 2-6 players. (you and 1 up to 5 AI bots)
Cards: A deck of cards with numbers 1-6 and llamas, with eight copies of each. 
Point Tokens: White (worth 1 point) and black (worth 10 points) are used for scoring. 
Dealing: Each player gets six cards dealt face down. 
Discard Pile: One card is flipped from the deck to start the discard pile. 
# Gameplay:
1. Turns:
Players take turns in clockwise order. 
2. Playing a Card:
On your turn, you can play a card from your hand if it matches the top card of the discard pile or is one number higher. Llamas can be played on 6s or other llamas. 
3. Drawing a Card:
If you can't or don't want to play, you can draw the top card from the deck. 
4. Folding:
If you choose to fold, you place your remaining cards face down and take no further action in the round. 
5. Round End:
The round ends when one player empties their hand or all players have folded. 

Scoring:
Cards in Hand: Each unique number card in your hand is worth its face value in points (e.g., a 3 is worth 3 points). 
Llamas: Each llama card in your hand is worth 10 points. 
Discarded Cards: Cards played and discarded are not counted towards your score. 
Returning Tokens: If you manage to play all your cards, you can return one token (black or white) to the supply. 
Winning:
Point Accumulation: Players accumulate points based on the cards left in their hand at the end of each round. 
Game End: The game ends when at least one player has 40 or more points. 
Winner: The player with the lowest score wins. 



