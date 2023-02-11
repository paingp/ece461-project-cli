# ece461-project-cli

This custom command line interface application takes in a file of npmjs or github urls
to analyze, rank, and sort based on a variety of factors.

####Ranking Factors
1) Correctness: How correct the module’s outputs are for Sarah’s needs 
2) Bus Factor: Measurement of risk for continued maintenance of the project
3) Responsive Maintainer: How responsive maintainers of the module are
4) Ramp-up Time: How easy it is for engineers to learn to use the module
5) License Compatibility: Compatibility with GNU Lesser General Public License v2.1 (LGPLv2.1)
6) Net Scoe: Weighted sum of first five factors

####Run Commands
1) ./run install: Installs the dependencies and returns exit 0 on success or 1 on failure
2) ./run build: Completes any compilation and returns exit 0 on success or 1 on failure
3) ./run URL_FILE: Takes in a text file of URLS and scores an analyzes them for each of the factors above
4) ./run test: runs a test suite, tests passed/tests total, coverage and returns exit 0 on success or 1 on failures

####How To Run Application
Simply just type in any of the commands listed under "Run Commands" section
