export function DeleteButton({ onClick }: { onClick: () => void }) {
  return (
    <button className="btn" onClick={onClick}>
      <span className="mx-1">x</span>
    </button>
  );
}
