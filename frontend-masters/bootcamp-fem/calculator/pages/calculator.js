let runningTotal = 0
let windowDisplay = '0'
let previousOperator = null

const screen = document.querySelector('.screen')

function buttonClick(value) {
  if (isNaN(value)) {
    handleSymbol(value)
  } else {
    handleNumber(value)
  }
  screen.innerText = windowDisplay
}

function handleSymbol(symbol) {
  console.log(symbol)
  // If equal is pressed
  switch (symbol) {
    case 'C':
      windowDisplay = 0
      runningTotal = 0
      break
    case '←':
      if (windowDisplay.length === 1) {
        windowDisplay = '0'
      } else {
        windowDisplay = windowDisplay.substring(0, windowDisplay.length - 1)
      }
      break
    case '=':
      if (previousOperator === null) {
        return
      }
      flushOperation(parseInt(windowDisplay))
      previousOperator = null
      windowDisplay = runningTotal
      runningTotal = 0
      break
    case '+':
    case '−':
    case '×':
    case '÷':
      handleMath(symbol)
      break
  }
  // If another operation is pressed
}

function handleMath(symbol) {
  if (windowDisplay === '0') {
    return
  }

  const intBuffer = parseInt(windowDisplay)

  if (runningTotal === 0) {
    runningTotal = intBuffer
  } else {
    flushOperation(intBuffer)
  }

  previousOperator = symbol
  windowDisplay = '0'
}

function flushOperation(intBuffer) {
  switch (previousOperator) {
    case '+':
      runningTotal += intBuffer
      break
    case '−':
      runningTotal -= intBuffer
      break
    case '×':
      runningTotal *= intBuffer
      break
    case '÷':
      runningTotal /= intBuffer
      break
  }
}

function handleNumber(numberString) {
  if (windowDisplay === '0') {
    windowDisplay = numberString
  } else {
    windowDisplay += numberString
  }
}

function init() {
  document
    .querySelector('.calc-buttons')
    .addEventListener('click', function (event) {
      buttonClick(event.target.innerText)
    })
}

init()
