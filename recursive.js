function fac (n) {
    if (n === 0|| n===1) return n
    else {
        return fac(n-1)+fac(n-2)
    }
}


console.time('fac')
console.log(fac(20))
console.timeEnd('fac')