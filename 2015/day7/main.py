from itertools import product

def charger_donnees(fichier):
    """
    Charge les données depuis un fichier texte.
    Chaque ligne doit être au format :
    test_value: number1 number2 ...
    
    Args:
    fichier (str): Chemin vers le fichier texte contenant les données.

    Returns:
    list: Une liste de tuples (test_value, [numbers]).
    """
    equations = []
    
    with open(fichier, 'r') as f:
        for ligne in f:
            # Supprimer les espaces inutiles
            ligne = ligne.strip()
            if not ligne:
                continue  # Sauter les lignes vides
            
            # Séparer le test_value des nombres
            test_value, numbers = ligne.split(":")
            test_value = int(test_value.strip())
            numbers = list(map(int, numbers.strip().split()))
            
            equations.append((test_value, numbers))
    
    return equations

def left_to_right_eval(numbers, operators):
    """
    Evaluate a sequence of numbers with given operators left-to-right.
    """
    result = numbers[0]
    for i, op in enumerate(operators):
        if op == "+":
            result += numbers[i + 1]
        elif op == "*":
            result *= numbers[i + 1]
        elif op == "||":
            result = int(str(result) + str(numbers[i + 1]))
    return result

def process_equations(equations):
    """
    Process a list of equations to find the total calibration result.
    
    Args:
    equations (list): A list of tuples (test_value, [numbers]).
    
    Returns:
    int: The total calibration result.
    """
    total_calibration = 0
    
    for test_value, numbers in equations:
        n = len(numbers)
        valid = False
        
        # Generate all possible operator combinations
        for ops in product(["+", "*", "||"], repeat=n - 1):
            if left_to_right_eval(numbers, ops) == test_value:
                valid = True
                break
        
        if valid:
            total_calibration += test_value
    
    return total_calibration

# Charger les données depuis un fichier texte
chemin_fichier = "input.txt"
equations = charger_donnees(chemin_fichier)

# Vérifier les données chargées
print("Équations chargées :")
print(equations)

# Calculer le résultat total des calibrations
resultat_total = process_equations(equations)
print(f"Résultat total des calibrations : {resultat_total}")
