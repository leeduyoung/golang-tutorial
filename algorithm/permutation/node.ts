function permutation(arr: number[], selectNum: number): number[][] {
  let ret: number[][] = []

  if (selectNum === 1) return arr.map((item) => [item])

  for (let i = 0; i < arr.length; i++) {
    const fixer = arr[i]

    let restArr = []
    for (const t of arr) if (t !== arr[i]) restArr.push(t)

    const result = permutation(restArr, selectNum - 1)
    for (let j = 0; j < result.length; j++) {
      const data = [fixer, ...result[j]]
      ret.push(data)
    }
  }

  return ret
}

console.log(permutation([1, 2, 3, 4], 3)) // 24가지
console.log(permutation([1, 2, 3, 4], 3).length) // 24가지
