function isPrime(num) {
  if (num % 2 === 0) {
    return false
  }
  for (var i = 3; i < num; i += 2) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function generatePrimes(n) {
  var arr = [2];
  for (var i = 3; i >= 3; i += 2) {
    if (isPrime(i)) {
      arr.push(i);
      if (arr.length === n) {
        return arr
      }
    }
  }
  return arr
}


const myArgs = process.argv.slice(2);
let primes = generatePrimes(parseInt(myArgs[0]))

console.log(`Generated ${primes.length} primes`)