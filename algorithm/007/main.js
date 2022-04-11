function solution(bridge_length, weight, truck_weights) {
  let time = 0
  let sum = 0
  let bridgeQueue = new Array(bridge_length).fill(0)
  let completeQueue = []
  const truckCount = truck_weights.length

  while (completeQueue.length !== truckCount) {
    const firstItem = bridgeQueue.shift()
    sum -=  firstItem

    if (truck_weights.length) {
      if (sum + truck_weights[0] > weight) {
        bridgeQueue.push(0)
      } else {
        const firstTruck = truck_weights.shift()
        bridgeQueue.push(firstTruck)
        sum += firstTruck
      }
    }

    if (firstItem !== 0) {
      completeQueue.push(firstItem)
    }

    time++
  }

  return time
}

console.log(solution(2, 10, [7,4,5,6])) // 8
console.log(solution(100, 100, [10])) // 101