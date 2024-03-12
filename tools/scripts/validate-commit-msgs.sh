#!/bin/bash

commitSummaryRegex="^(\+|\-)\s[a-z0-9]+\s(.+)$"
commitTypeRegex="^(chore|docs|feat|fix|perf|ref|revert|style|test|wip)"
commitScopeRegex="^[a-z]+(\((a-rule|a-schema|acl|app-scope|bot-cfg|bot-rule|bot|c-rule|cache|cdn|ci-cd|client|common|config|deployment|e-rule|env-var|env|internal|m-role|org|property|r-rule|tls|token|util)\))?\:\s"
issueNumberRegex="\s(\[#[0-9]+\]|\(#[0-9]+\))$"
defectiveCommitSummaries=()

while IFS= read -r line; do
    [[ $line =~ $commitSummaryRegex ]]
    commitSummary="${BASH_REMATCH[2]}"

    if [ ${#commitSummary} -ge 72 ]; then
        defectiveCommitSummaries+=("Commit summary \`$commitSummary\` is more than 72 characters long")
    fi

    if [[ !($commitSummary =~ $commitTypeRegex) ]]; then
        defectiveCommitSummaries+=("Commit summary \`$commitSummary\` does not respect commit type restrictions")
    fi

    if [[ !($commitSummary =~ $commitScopeRegex) ]]; then
        defectiveCommitSummaries+=("Commit summary \`$commitSummary\` does not respect commit scope restrictions")
    fi

    if [[ !($commitSummary =~ $issueNumberRegex) ]]; then
        defectiveCommitSummaries+=("Commit summary \`$commitSummary\` does include an issue number")
    fi
done <<< "$(git cherry -v origin/beta)"

echo "### Commit messages check result" >> $GITHUB_STEP_SUMMARY

if [ ${#defectiveCommitSummaries[@]} -eq 0 ]; then
    echo "All commit messages are valid" >> $GITHUB_STEP_SUMMARY
    exit 0
else
    echo "There are ${#defectiveCommitSummaries[@]} uncompliant commit messages" >> $GITHUB_STEP_SUMMARY
    echo "#### Checks performed" >> $GITHUB_STEP_SUMMARY

    echo "- Commit summary size: 72 characters maximum" >> $GITHUB_STEP_SUMMARY
    echo "- Commit type constraints: \`/$commitTypeRegex/\`" >> $GITHUB_STEP_SUMMARY
    echo "- Commit scope constraints: \`/$commitScopeRegex/\`" >> $GITHUB_STEP_SUMMARY
    echo "- Issue number presence: \`/$issueNumberRegex/\`" >> $GITHUB_STEP_SUMMARY

    echo "#### Invalid commit messages" >> $GITHUB_STEP_SUMMARY

    for invalid in "${defectiveCommitSummaries[@]}"; do
        echo "- $invalid" >> $GITHUB_STEP_SUMMARY
    done

    echo "" >> $GITHUB_STEP_SUMMARY
    echo "" >> $GITHUB_STEP_SUMMARY
    echo "Please, rebase your branch fixing the problematic commit summaries and update your PR :)" >> $GITHUB_STEP_SUMMARY

    exit 1
fi
