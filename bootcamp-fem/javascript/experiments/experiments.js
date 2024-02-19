const dog = {
  name: 'dog',
  speak() {
    console.log('WOOF')
  },
}

dog.speak()

function makeDogSpeak() {
  return dog.speak()
}

// describe('makeDogSpeak', () => {
//   if('makes dog speak', () => {
//
// })
