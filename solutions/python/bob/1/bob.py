def response(hey_bob):
    cleaned = hey_bob.strip()
    if not cleaned:
        return "Fine. Be that way!"
    
    is_question = cleaned.endswith("?")
    is_yelling = cleaned.isupper()
    
    if is_yelling and is_question:
        return "Calm down, I know what I'm doing!"
    if is_yelling:
        return "Whoa, chill out!"
    if is_question:
        return "Sure."
    
    return "Whatever."