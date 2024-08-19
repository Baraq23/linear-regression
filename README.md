# linear-stats

This program reads data from a txt file which contains data points at each line and outputs the `linear regression line` and the `Pearson correlation coefficient`. The positions of the datapoint will be the x-values and the data points will is the y-values


## Usage

This program only runs in 64bit operating systems. The data provided must be in a txt file named data.txt and must contain values in each line of the data file.


## Installation & usage

To run the program,follow these steps:

* clone the program into your computer

```bash
https://barraotieno/gitea.com/linear-stats
```

* Open the cloned folder with your prefered text editor among VScode and Vim.

* Replace the content of the data.txt with your own values to be computed

* In the terminal from your text editor type the following:
```bash
go run . data.txt
```

* Hit Enter. the Results will be shown in the terminal.

## Output

If the data provided in the data.txt file are, for example:

```bash
18 22 25 29 24 13 61
```

The output will be as follows:

```bash
Linear Regression line: y = 3.928571x + 11.714286
Pearson Correlation Coefficient: 0.5415106754
```

## Notion


### Linear Regression Line
![](/imgs/linear-regression-equation-1.png)

### Pearson's Correlation Coefficient

*Correlation:* It measures the relationship or association between two variables. Pearson's correlation specifically looks at how well a straight line can describe the relationship between the variables.
Pearson's r: This value ranges from -1 to 1:

![](/imgs/correlation-visualized.png)

- +1: A perfect positive correlation. As one variable increases, the other variable also increases in a perfectly straight line.
- 0: No correlation. The variables do not have any linear relationship.
- -1: A perfect negative correlation. As one variable increases, the other decreases in a perfectly straight line.

_Example:_

- *Positive Correlation:* If you look at the relationship between the number of hours studied and exam scores, you might find a positive correlation, meaning more study hours generally lead to higher scores.
- *Negative Correlation:* If you look at the relationship between the number of hours spent watching TV and exam scores, you might find a negative correlation, meaning more TV watching is associated with lower scores.

#### Calculations of Pearson's Correlation Coefficient

![](/imgs/correlation-formula.png)
![](/imgs/correlation-formula-interpretation.png)


## Author
- [Barrack otieno](https://www.linkedin.com/in/barrack-kope-064a43244).