flow "Del diálogo al artefacto" {
  step conversation "Conversación" {
    text "Exploración y discusión"
    icon "message"
  }

  split {
    step journey "Journey" {
      text "Cómo llegamos"
      icon "route"
    }
    step edr "EDR" {
      text "Qué decidimos"
      icon "decision"
    }
  }

  step knowledge "Conocimiento" {
    icon "book"
  }
  step artifacts "Artefactos" {
    icon "boxes"
  }

  highlight "La experiencia y la decisión se convierten en conocimiento reutilizable"
}
