function combination(arr: number[], selectNum: number): number[][] {
  let ret: number[][] = []

  if (selectNum === 1) return arr.map((item) => [item])

  for (let i = 0; i < arr.length; i++) {
    const fixer = arr[i]
    const restArr = arr.slice(i + 1)
    const result = combination(restArr, selectNum - 1)

    for (let j = 0; j < result.length; j++) {
      const data = [fixer, ...result[j]]
      ret.push(data)
    }
  }

  return ret
}

console.log(combination([1, 2, 3, 4], 3)) // 4가지
console.log(combination([1, 2, 3, 4], 3).length) // 4가지

// ==============================================================================

console.log(combination2([1, 2, 3, 4], 3)) // 4가지
console.log(combination2([1, 2, 3, 4], 3).length) // 4가지

function combination2(arr: number[], selectNum: number): number[][] {
  let ret: number[][] = []

  for (let i = 0; i < arr.length; i++) {
    for (let j = i + 1; j < arr.length; j++) {
      for (let k = j + 1; k < arr.length; k++) {
        if (i === j || j === k || k === i) continue

        ret.push([arr[i], arr[j], arr[k]])
      }
    }
  }

  return ret
}
