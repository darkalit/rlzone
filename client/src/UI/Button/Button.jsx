import React from "react";
import "./Button.style.css";

export const ButtonType = Object.freeze({
  Fill: "fill",
  Outline: "outline",
  Gradient: "gradient",
  GradientNoGlow: "gradient-no-glow",
  Ban: "ban",
});

export default function Button({ pv = "8px", ph = "24px", width, type, text }) {
  const style = {
    padding: `${pv} ${ph}`,
    whiteSpace: "nowrap",
  };

  if (width) {
    style.width = width;
  }

  let typeClass = null;

  switch (type) {
    case ButtonType.Fill:
      typeClass = "button-fill";
      break;
    case ButtonType.Outline:
      typeClass = "button-outline";
      break;
    case ButtonType.Gradient:
      typeClass = "button-gradient";
      break;
    case ButtonType.GradientNoGlow:
      typeClass = "button-gradient-no-glow";
      break;
    case ButtonType.Ban:
      typeClass = "button-ban";
      break;
    default:
      return null;
  }

  return (
    <button className={typeClass} style={style}>
      {text}
    </button>
  );
}
