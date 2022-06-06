function solution(name) {
  var answer = 0
  const length = name.length
  let upDownCount = 0
  let leftRightCountList = [length - 1] //한 방향으로 쭉 갔을 때
  for (let i = 0; i < length; i++) upDownCount += minUpOrDownCount(name[i])
  for (let startOfA = 0; startOfA < name.length; startOfA++) {
    let endOfA = startOfA + 1
    while (endOfA < length && name[endOfA] === 'A') endOfA++
    const [moveToStartOfA, moveToEndOfA] = [startOfA, length - endOfA]

    // 0 -> A.., 0 <- A.., ..A <- -1
    leftRightCountList.push(moveToStartOfA * 2 + moveToEndOfA)

    //시작부터 뒤로 가는 경우 ..A <- -1, ..A -> -1, 0 -> A..
    leftRightCountList.push(moveToEndOfA * 2 + moveToStartOfA)
  }
  answer = upDownCount + Math.min(...leftRightCountList)
  return answer
}

function minUpOrDownCount(destination) {
  const Alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  const index = Alphabet.indexOf(destination)
  return Math.min(index, Alphabet.length - index)
}

console.log(solution('JEROEN')) // 56
console.log(solution('JAN')) // 23
