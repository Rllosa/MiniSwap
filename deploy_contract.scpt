tell application "iTerm2"
    tell current session of current window
        set currentDir to (do shell script "pwd")
    end tell
    tell current window
        create tab with default profile
        set newTab to current tab
        tell current session of newTab
            write text "cd " & currentDir & " && ganache --gasLimit 120000000 --chain.chainId 11155420 > ganache-output.txt"
        end tell
    end tell
    -- Switch back to the original tab
    tell first tab of current window
        select
    end tell
end tell