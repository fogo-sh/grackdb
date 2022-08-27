export function DeleteButton({ onClick }: { onClick?: () => void }) {
  return (
    <button className="btn" onClick={onClick}>
      <p className="mx-1">x</p>
    </button>
  );
}
