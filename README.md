# NPM Module Rater

This command line interface application takes in a file containing URLs to NPM packages or GitHub repositories and
analyzes how useful the open source package associated with each URL is based on a number of factors.

#### Ranking Factors
1) Correctness: How correct the module’s outputs are for ACME corporation's needs 
2) Bus Factor: Measurement of risk for continued maintenance of the project
3) Responsive Maintainer: How responsive maintainers of the module are
4) Ramp-up Time: How easy it is for engineers to learn to use the module
5) License Compatibility: Compatibility with GNU Lesser General Public License v2.1 (LGPLv2.1)
6) Net Score: Weighted sum of first five factors

#### Run Commands
1) ./run install: Installs the dependencies and returns exit 0 on success or 1 on failure
2) ./run build: Completes any compilation and returns exit 0 on success or 1 on failure
3) ./run URL_FILE: Takes in a text file of URLS and scores an analyzes them for each of the factors above
4) ./run test: runs a test suite, tests passed/tests total, coverage and returns exit 0 on success or 1 on failures

#### Requirement To Run Application
Our system requires following environment variables to be provided:
GITHUB_TOKEN: GitHub personal access token
LOG_LEVEL: (0 means silent, 1 means informational messages, 2 means debug messages)
LOG_FILE: filepath to store logs

#### Citations
1) https://stackoverflow.com/questions/71153302/how-to-set-depth-for-recursive-iteration-of-directories-in-filepath-walk-func
2) https://rakaar.github.io/posts/2021-04-23-go-json-res-parse/
3) https://stackoverflow.com/questions/24809287/how-do-you-get-a-golang-program-to-print-the-line-number-of-the-error-it-just-ca

Project Instructions
CLI for Trustworthy Modules Project Description.pdf

Team Members:\
Paing Khant\
Matthew Campbell\
Aditya Srikanth\
Brian Ng
