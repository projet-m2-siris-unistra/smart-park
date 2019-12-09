from sanic_wtf import SanicForm
from wtforms import PasswordField, StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit



# Creation form
class CreationForm(SanicForm):
    name = StringField(
        widget=BXInput(input_type="text"),
        label='Nom',
        description='Ce nom unique identifiera la zone.',
        validators=[DataRequired(), Length(max=40)]
    )

    type = SelectField(
        widget=BXSelect(),
        label='Type de la zone',
        description='Indiquer le type de parking que contiendra la zone.',
        validators=[DataRequired()],
        choices=[
            ('free', 'Gratuit'),
            ('paid', 'Payant'),
            ('blue', 'Zone bleue')
        ]
    )

    color = StringField(
        widget=BXInput(input_type="color"),
        label='Couleur',
        description='Choisissez une ,couleur associée à la zone. Cela facilitera sa reconnaissance sur les cartes.',
        validators=[DataRequired()]
    )

    polygon = HiddenField()

    submit = SubmitField(
        widget=BXSubmit(),
        label='Suivant'
    )