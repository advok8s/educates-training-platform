# Generated by Django 4.2.8 on 2023-12-22 02:48

from django.db import migrations
import project.apps.workshops.models


class Migration(migrations.Migration):

    dependencies = [
        ("workshops", "0009_alter_trainingportal_generation_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="trainingportal",
            name="labels",
            field=project.apps.workshops.models.JSONField(default={}),
        ),
    ]