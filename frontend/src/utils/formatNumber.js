export function formatNumber(value, minimumFractionDigits = 2) {
  return Number(value).toLocaleString("en-US", {
    minimumFractionDigits,
  });
}
