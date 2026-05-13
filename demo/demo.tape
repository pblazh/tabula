# Install VHS:
#   brew install vhs

Output ../demo.gif

Require tabula

Set Theme "GruvboxDark"
Set FontSize 22
Set FontFamily "FiraCode Nerd Font Mono"
Set Width 860
Set Height 480
Set Padding 24
Set TypingSpeed 55ms
Set PlaybackSpeed 1.0
Set Shell "bash"

# Reset demo state silently before we start recording visible frames
Hide
Type "cp input.template.csv input.csv && nvim input.csv"
Enter
Sleep 300ms
Show

# Recording
Sleep 1800ms
Type 'ggg^fs'
Sleep 400ms
Type 'gf'
Sleep 1800ms
Ctrl+o
Sleep 400ms
Type 'g^jjj'
Sleep 1800ms
Type 'f1cw22'
Escape
Sleep 1800ms
Type 'kcw3'
Escape
Type ' fs'
Sleep 2800ms
