from sanic_wtf import SanicForm
from wtforms import PasswordField, StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit


class IssueForm(SanicForm):
    type = SelectField(
        widget=BXSelect(),
        label='Type de problème',
        description='Indiquer le type de problème auquel vous êtes confrontés.',
        validators=[DataRequired()],
        choices=[
            ('1', 'Technique'),
            ('2', 'Demande particulière'),
            ('3', 'Renseignements'),
            ('4', 'Autre')
        ]
    )
    
    report = StringField(
        widget=BXInput(input_type="text"),
        label="",
        description="Veuillez décrire précisement votre problème.",
        validators=[DataRequired(), Length(max=500)]
    )

    submit = SubmitField(
        widget=BXSubmit(),
        label='Envoyer'
    )
