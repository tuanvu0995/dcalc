import { useCallback } from "react";
import { useAppState } from "../../context/AppStateContext";
import { formatNumber } from "../../utils/formatNumber";

export default function LineItem({ node, index }) {
  const { state } = useAppState();

  const callback = useCallback(() => node.buffer.split(""), [node.buffer]);

  const renderWord = () => {
    const words = callback();
    const lastIndex = words.length - 1;
    return words.map((word, wordIndex) => (
      <span
        key={wordIndex + "-" + word.kind}
        className={
          state.nodeIndex === index && wordIndex === lastIndex ? "cursor" : ""
        }
      >
        {"+-*/".includes(word) ? <>&nbsp;{word}&nbsp;</> : word}
      </span>
    ));
  };

  return (
    <tr>
      <td className="line-item__name" width={20}>
        <span className="tag">{node.operator}</span>
      </td>
      <td className="line-item__value">
        <span className={state.nodeIndex === index ? "cursor" : ""}>
          {renderWord()}
        </span>
      </td>
      <td className="line-item__result" align="right">
        <span>{formatNumber(node.result)}</span>
      </td>
    </tr>
  );
}
