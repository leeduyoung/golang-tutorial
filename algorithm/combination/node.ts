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
