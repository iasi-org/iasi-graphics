package theme

type Theme struct {
	Background, Card, CardStroke, Text, Muted, Accent, AccentSoft, Connector string
	FontFamily                                                               string
}

func IASI() Theme {
	return Theme{Background: "#F7F8FC", Card: "#FFFFFF", CardStroke: "#DDE3EE", Text: "#172033", Muted: "#58657A", Accent: "#6C55E0", AccentSoft: "#EEEAFE", Connector: "#9AA6B8", FontFamily: "Inter, Segoe UI, Arial, sans-serif"}
}
