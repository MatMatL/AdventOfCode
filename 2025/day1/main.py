def main():
    # Read the input file
    with open("./test.txt", "r") as f:
        lines = f.readlines()

    # Process the lines
    # for line in lines:
    #     print(line.strip())

    result = SolvePart1(lines)

    print("Result:", result)

def SolvePart1(lines):
    result = 0
    pointingAt = 50

    for line in lines:
        direction = line[0]
        amount = int(line[1:])

        expectedPointing = pointingAt
        changed = False

        if direction == "L":
            expectedPointing -= amount
        elif direction == "R":
            expectedPointing += amount
        
        if expectedPointing < 0:
            times = abs(expectedPointing/100)
            changed = True
            if (times > int(times)):
                result += int(times)+1
            else: 
                result +=1
            
            expectedPointing += 100 * int(times)
            pointingAt = 100 - abs(expectedPointing)
        
        elif expectedPointing >= 100:
            times = int(expectedPointing/100)
            changed = True
            result += times
            pointingAt = expectedPointing - (100 * times)
        
        else:
            if direction == "L":
                pointingAt -= amount
            elif direction == "R":
                pointingAt += amount

        print("From line: ", line.strip(), "-> Pointing at: ", pointingAt)
        
        if changed:
            print("New result ! : ", result)


    return result


if __name__ == "__main__":
    main()