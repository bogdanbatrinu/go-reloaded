
func replaceHex(s string) string {
    re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s+\(hex\)`)
    return re.ReplaceAllStringFunc(s, func(match string) string {
        parts := re.FindStringSubmatch(match)
        if parts == nil {
            return match
        }
        v, err := strconv.ParseInt(parts[1], 16, 64)
        if err != nil {
            return match
        }
        return strconv.FormatInt(v, 10)
    })
}

func replaceBin(s string) string {
    re := regexp.MustCompile(`\b([01]+)\s+\(bin\)`)
    return re.ReplaceAllStringFunc(s, func(match string) string {
        parts := re.FindStringSubmatch(match)
        if parts == nil {
            return match
        }
        v, err := strconv.ParseInt(parts[1], 2, 64)
        if err != nil {
            return match
        }
        return strconv.FormatInt(v, 10)
    })
}

func applyCaseModifiers(s string) string {

    s = regexp.MustCompile(`(\S+)\s+\(up\)`).ReplaceAllStringFunc(s, func(m string) string {
        sub := regexp.MustCompile(`^(\S+)\s+\(up\)$`).FindStringSubmatch(m)[1]
        return strings.ToUpper(sub)
    })
    s = regexp.MustCompile(`(\S+)\s+\(low\)`).
        ReplaceAllStringFunc(s, func(m string) string {
            sub := regexp.MustCompile(`^(\S+)\s+\(low\)$`).FindStringSubmatch(m)[1]
            return strings.ToLower(sub)
        })
    s = regexp.MustCompile(`(\S+)\s+\(cap\)`).
        ReplaceAllStringFunc(s, func(m string) string {
            sub := regexp.MustCompile(`^(\S+)\s+\(cap\)$`).FindStringSubmatch(m)[1]
            return strings.Title(strings.ToLower(sub))
        })

    reN := regexp.MustCompile(`\(\s*(up|low|cap),\s*(\d+)\s*\)`)
    return reN.ReplaceAllStringFunc(s, func(m string) string {
        parts := reN.FindStringSubmatch(m)
        mode, count := parts[1], parts[2]
        n, _ := strconv.Atoi(count)
        tokens := strings.Fields(s[:strings.Index(s, m)])
        before := tokens[len(tokens)-n:]
        var transformed []string
        for _, w := range before {
            switch mode {
            case "up":
                transformed = append(transformed, strings.ToUpper(w))
            case "low":
                transformed = append(transformed, strings.ToLower(w))
            case "cap":
                transformed = append(transformed, strings.Title(strings.ToLower(w)))
            }
        }
        prefix := strings.Join(tokens[:len(tokens)-n], " ")
        suffix := s[strings.Index(s, m)+len(m):]
        return prefix + " " + strings.Join(transformed, " ") + suffix
    })
}

func fixPunctuation(s string) string {

    s = strings.ReplaceAll(s, " ...", "...")
    s = strings.ReplaceAll(s, "... ", "...")

    s = strings.ReplaceAll(s, " !?", "!?")
    s = strings.ReplaceAll(s, "?! ", "?!")

    re := regexp.MustCompile(`\s*([,\.!\?:;]+)(\S)`)
    s = re.ReplaceAllString(s, "$1 $2")

    s = regexp.MustCompile(`\s+([,\.!\?:;]+)`).ReplaceAllString(s, "$1")
    return s
}


func fixSingleQuotes(s string) string {
    re := regexp.MustCompile(`'\s*([^']+?)\s*'`)
    return re.ReplaceAllString(s, `'$1'`)
}

func fixArticles(s string) string {
    re := regexp.MustCompile(`\b([Aa])\s+([AaEeIiOoUuHh]\w*)`)
    return re.ReplaceAllString(s, "an $2")
}
