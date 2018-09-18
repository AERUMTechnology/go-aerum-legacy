#!/bin/bash
########################################################################
########################################################################
############# Terminal Launcher for Aerum Consensus DevKit #############
###################### Supports OSX and Linux ##########################
################### Launched 7 terminals at once #######################
########################################################################
########################################################################


platform='unknown'

unamestr=`uname`

if [[ "$unamestr" == 'Linux' ]]; then
   platform='linux'
elif [[ "$unamestr" == 'Darwin' ]]; then
   platform='osx'
fi

if [[ $platform == 'linux' ]]; then

gnome-terminal  --window --working-directory=$PWD \
                --tab -e "bash $PWD/start-scripts/node1.sh" \
                --tab -e "bash $PWD/start-scripts/node2.sh" \
                --tab -e "bash $PWD/start-scripts/node3.sh" \
                --tab -e "bash $PWD/start-scripts/node4.sh" \
                --tab -e "bash $PWD/start-scripts/node5.sh" \
                --tab -e "bash $PWD/start-scripts/node6.sh" \
                --tab -e "bash $PWD/start-scripts/node7.sh"
    
elif [[ $platform == 'osx' ]]; then

osascript -e "tell application \"Terminal\"" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node1.sh\" in front window" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node2.sh\" in front window" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node3.sh\" in front window" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node4.sh\" in front window" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node5.sh\" in front window" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node6.sh\" in front window" \
          -e "tell application \"System Events\" to keystroke \"t\" using {command down}" \
          -e "do script \"bash $PWD/start-scripts/node7.sh\" in front window" \
          -e "end tell" > /dev/null   
         
fi