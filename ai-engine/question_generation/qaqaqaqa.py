from pipelines import pipeline

nlp = pipeline("question-generation")
while True:
    text = input("Enter text: ")
    print(nlp(text))
