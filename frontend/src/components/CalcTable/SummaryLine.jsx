import { useAppState } from "../../context/AppStateContext";
import { formatNumber } from "../../utils/formatNumber";

export default function SummaryLine() {
  const { state } = useAppState();

  return (
    <div className="summary-tr">
      <div className="line-item__result" align="right">
        <span>= {formatNumber(state.value)}</span>
      </div>
    </div>
  );
}
