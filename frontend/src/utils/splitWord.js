// word = "(5*10)-2/2" || "min(1, 10)"
export default function splitWord(word) {
  const operators = ["+", "-", "*", "/", "(", ")"];
  const functions = ["min", "max", "avg"];
  const result = [];
  let temp = "";
  for (let i = 0; i < word.length; i++) {
    if (operators.includes(word[i])) {
      if (temp) {
        result.push(temp);
        temp = "";
      }
      result.push(word[i]);
    } else if (word[i] === " ") {
      if (temp) {
        result.push(temp);
        temp = "";
      }
    } else {
      temp += word[i];
    }
  }
  if (temp) {
    result.push(temp);
  }
  return result;
}
