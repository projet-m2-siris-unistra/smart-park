from sanic_wtf import SanicForm
from wtforms import StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit



class BaseForm(SanicForm):
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
        render_kw={'class' : ''},
        label='Couleur',
        description='Choisissez une ,couleur associée à la zone. Cela facilitera sa reconnaissance sur les cartes.',
        validators=[DataRequired()]
    )



# Creation form
class CreationForm(BaseForm):
    polygon = HiddenField()

    submit = SubmitField(
        widget=BXSubmit(),
        label='Suivant'
    )


# Configuration form
class ConfigurationForm(CreationForm):
    submit = SubmitField(
        widget=BXSubmit(),
        label='Enregistrer'
    )

    delete = SubmitField(
        widget=BXSubmit(),
        render_kw={'color': 'danger'},
        label="Supprimer"
    )